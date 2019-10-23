/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var bucketLead = []byte("leads")

type Lead struct {
	Visitor
	Email string `json:"email"`
}

func addLead(lead Lead) (*Lead, error) {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketLead)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		lead.ID = id
		buf, err := data.Encode(lead)
		if err != nil {
			return err
		}
		return b.Put(data.IntToByteArray(lead.ID), buf)
	})
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

func getLead(id uint64) (*Lead, error) {
	var buf []byte

	err := data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketLead)

		buf = b.Get(data.IntToByteArray(id))
		return nil
	})

	if len(buf) == 0 {
		return nil, nil
	}

	var l Lead
	if err := data.Decode(buf, &l); err != nil {
		return nil, err
	}

	return &l, err
}

func updateLead(id uint64, l Lead) {
}

func removeLead(id uint64) error {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketLead)

		return b.Delete(data.IntToByteArray(id))
	})
	return err
}
