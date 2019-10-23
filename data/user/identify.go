/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var bucketIdentify = []byte("identity")

type Identification struct {
	ReferenceID     uint64
	ConnectionToken string
	IsVisitor       bool
	IsLead          bool
	IsUser          bool
}

// Identify determines if an incoming connection is a visitor, a lead or a user.
//
// The identity bucket can be seen as the sessions between the browser and
// the underlying entity.
//
// The flow is as follow:
// 1. Look into the identity bucket to see if there's a match for this token
// 2. If there's no match:
// -  a. If there's an email supplied we add as a user
// -  b. otherwise we add as a visitor
// 3. If there's a match, we update the proper entity
//
// The websocket connection will need to adjust its token (client and service side)
// once the identity has been decided.
//
// The flow for a typical user identification:
// 1. User might visit marketing website www.product.com
// - they are tagged as a visitor at that moment with a fresh token
// 2. If they visit the login page app.product.com/login
// - The identity will be called with the cache token, and this new page load
// - connection will be tagged as the visitor from step #1
// 3. When they submit the authentication and the server authenticate them:
// - The call to Identify will now have the real User data.
// - The session will be matched from the cache token, and will be promoted
// - from visitor/lead to user.
// - The client cache token and the server client token will need to be updated
// - to the new user token.
func Identify(token string, msg string) (*Identification, error) {
	var id Identification
	var user User
	if err := json.Unmarshal([]byte(msg), &user); err != nil {
		return nil, fmt.Errorf("unable to parse message as a User: %v", err)
	}

	// if there's a supplied connection token we take that one
	if len(user.ConnectionToken) > 0 {
		token = user.ConnectionToken
	}

	var buf []byte
	data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketIdentify)

		buf = b.Get([]byte(token))
		return nil
	})
	if len(buf) == 0 {
		// if we've never seen them, we add them as visitor or user
		// depending if there's an email supplied

		if len(user.YourID) > 0 {
			u, err := addUser(user)
			if err != nil {
				return nil, fmt.Errorf("unable to add user: %v", err)
			}

			id.ReferenceID = u.ID
			id.ConnectionToken = token
			id.IsUser = true
		} else if len(user.Email) > 0 {
			l, err := addLead((user).Lead)
			if err != nil {
				return nil, fmt.Errorf("unable to add lead: %v", err)
			}

			id.ReferenceID = l.ID
			id.ConnectionToken = token
			id.IsLead = true
		} else {
			v, err := addVisitor((user).Visitor)
			if err != nil {
				return nil, fmt.Errorf("unable to add visitor: %v", err)
			}

			id.ReferenceID = v.ID
			id.ConnectionToken = token
			id.IsVisitor = true
		}

		err := updateIdentity(id)
		if err != nil {
			return nil, fmt.Errorf("unable to save the new identification: %v", err)
		}
		return &id, nil
	} // end of token not present in bucket

	// we've seen there before, let's update their info
	if err := data.Decode(buf, &id); err != nil {
		return nil, err
	}

	if id.IsVisitor {
		if len(user.YourID) > 0 {
			promoteToUser(&id, user)
		} else if len(user.Email) > 0 {
			promoteToLead(&id, (user).Lead)
		} else {
			go updateVisitor(id.ReferenceID, (user).Visitor)
		}
	} else if id.IsLead {
		if len(user.YourID) > 0 {
			promoteToUser(&id, user)
		} else {
			go updateLead(id.ReferenceID, (user).Lead)
		}
	} else if id.IsUser {
		go updateUser(id.ReferenceID, user)
	}
	return &id, nil
}

func updateIdentity(id Identification) error {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketIdentify)
		updated, err := data.Encode(id)
		if err != nil {
			return err
		}
		if err := b.Put([]byte(id.ConnectionToken), updated); err != nil {
			return err
		}
		return nil
	})
	return err
}

func promoteToLead(id *Identification, lead Lead) error {
	l, err := addLead(lead)
	if err != nil {
		return err
	}

	if err := removeVisitor(id.ReferenceID); err != nil {
		return err
	}

	id.ReferenceID = l.ID
	id.IsVisitor = false
	id.IsLead = true
	if err := updateIdentity(*id); err != nil {
		return err
	}

	return nil

}

func promoteToUser(id *Identification, user User) error {
	u, err := addUser(user)
	if err != nil {
		return err
	}

	if id.IsVisitor {
		if err := removeVisitor(id.ReferenceID); err != nil {
			return err
		}
	} else if id.IsLead {
		if err := removeLead(id.ReferenceID); err != nil {
			return err
		}
	}

	id.ReferenceID = u.ID
	id.IsVisitor = false
	id.IsLead = false
	id.IsUser = true
	if err := updateIdentity(*id); err != nil {
		return err
	}

	return nil
}
