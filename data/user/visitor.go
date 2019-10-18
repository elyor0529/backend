/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var bucketVisitor = []byte("visitors")

type Visitor struct {
	ID              uint64     `json:"id"`
	ConnectionToken string     `json:"connToken"`
	BrowserAgent    string     `json:"browserAgent"`
	SessionCount    int        `json:"sessionCount"`
	Trackings       []Tracking `json:"trackings"`
	Created         time.Time  `json:"created"`
	Updated         time.Time  `json:"updated"`
	LastActivity    time.Time  `json:"lastActivity"`
}

func addVisitor(v Visitor) (*Visitor, error) {
	v.Created = time.Now()
	v.Updated = v.Created
	v.LastActivity = v.Created
	v.SessionCount = 1

	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketVisitor)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		v.ID = id
		buf, err := data.Encode(v)
		if err != nil {
			return err
		}
		return b.Put(data.IntToByteArray(v.ID), buf)
	})
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func updateVisitor(id uint64, v Visitor) {
	var buf []byte
	data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketVisitor)
		buf = b.Get(data.IntToByteArray(id))
		return nil
	})
	if len(buf) == 0 {
		log.Printf("unable to find visitor %d for an update\n", id)
		return
	}

	var update Visitor
	if err := data.Decode(buf, &update); err != nil {
		log.Println(err)
		return
	}

	//TODO: make this better
	now := time.Now()
	if update.Updated.Day() != now.Day() {
		update.SessionCount++
	}

	update.Updated = now
	update.LastActivity = now
	update.BrowserAgent = v.BrowserAgent

	if len(v.Trackings) > 0 {
		update.Trackings = append(update.Trackings, v.Trackings[0])
	}

	buf, err := data.Encode(update)
	if err != nil {
		log.Println(err)
		return
	}

	err = data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketVisitor)
		return b.Put(data.IntToByteArray(id), buf)
	})
	if err != nil {
		log.Println(err)
	}
}

func removeVisitor(id uint64) error {
	return nil
}
