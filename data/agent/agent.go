/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package agent

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var bucketAgents = []byte("agents")

type Agent struct {
	ID       uint64    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:-`
	Token    string    `json:"token"`
	Created  time.Time `json:"created"`
}

func Add(a Agent) (*Agent, error) {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAgents)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		a.ID = id
		buf, err := data.Encode(a)
		if err != nil {
			return err
		}
		return b.Put([]byte(a.Email), buf)
	})
	return &a, err
}

func GetByEmail(email string) (*Agent, error) {
	var buf []byte
	err := data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAgents)

		buf = b.Get([]byte(email))
		return nil
	})
	if err != nil {
		return nil, err
	}

	var a Agent
	if err := data.Decode(buf, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func validateToken(email, token string) (*Agent, error) {
	a, err := GetByEmail(email)
	if err != nil {
		return nil, err
	} else if a.Token != token {
		return nil, fmt.Errorf("invalid email/token")
	}

	return a, nil
}
