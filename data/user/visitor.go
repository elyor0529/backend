/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"time"

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

func (v *Visitor) SetID(id uint64) {
	v.ID = id
}

func addVisitor(v Visitor) (*Visitor, error) {
	v.Created = time.Now()
	v.Updated = v.Created
	v.LastActivity = v.Created
	v.SessionCount = 1

	if err := data.CreateWithAutoIncrement(bucketVisitor, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func getVisitor(id uint64) (*Visitor, error) {
	var v Visitor
	if err := data.Get(bucketVisitor, data.IntToByteArray(id), &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func updateVisitor(id uint64, v Visitor) error {
	var cur Visitor
	if err := data.Get(bucketVisitor, data.IntToByteArray(id), &cur); err != nil {
		return err
	}

	//TODO: make this better
	now := time.Now()
	if cur.Updated.Day() != now.Day() {
		cur.SessionCount++
	}

	cur.Updated = now
	cur.LastActivity = now
	cur.BrowserAgent = v.BrowserAgent

	if len(v.Trackings) > 0 {
		cur.Trackings = append(cur.Trackings, v.Trackings[0])
	}

	return data.Update(bucketVisitor, data.IntToByteArray(id), cur)
}

func removeVisitor(id uint64) error {
	return data.Delete(bucketVisitor, data.IntToByteArray(id))
}
