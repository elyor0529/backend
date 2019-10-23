/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var bucketUser = []byte("users")

type User struct {
	Lead
	YourID         string            `json:"yourId"`
	FirstName      string            `json:"fname"`
	LastName       string            `json:"lname"`
	Avatar         string            `json:"avatar"`
	Tags           []string          `json:"tags"`
	Attributes     map[string]string `json:"customTags"`
	CompanyID      uint64            `json:"companyId"`
	CompanyName    string            `json:"companyName"`
	IsUnsubscribed bool              `json:"unsubscribed"`
}

type Tracking struct {
	Referrer string    `json:"referrer"`
	Source   string    `json:"utmSource"`
	Medium   string    `json:"utmMedium"`
	Campaign string    `json:"utmCampaign"`
	Terms    string    `json:"utmTerms"`
	Content  string    `json:"utmContent"`
	Tracked  time.Time `json:"tracked"`
}

func addUser(u User) (*User, error) {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketUser)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		u.ID = id
		buf, err := data.Encode(u)
		if err != nil {
			return err
		}
		return b.Put(data.IntToByteArray(u.ID), buf)
	})
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func getUser(id uint64) (*User, error) {
	var buf []byte

	err := data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketUser)

		buf = b.Get(data.IntToByteArray(id))
		return nil
	})

	if len(buf) == 0 {
		return nil, nil
	}

	var u User
	if err := data.Decode(buf, &u); err != nil {
		return nil, err
	}

	return &u, err
}

func updateUser(id uint64, u User) {

}

func removeUser(id uint64) error {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketUser)

		return b.Delete(data.IntToByteArray(id))
	})
	return err
}
