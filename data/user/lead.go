/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import (
	"github.com/parle-io/backend/data"
)

var bucketLead = []byte("leads")

type Lead struct {
	Visitor
	Email string `json:"email"`
}

func (l *Lead) SetID(id uint64) {
	l.ID = id
}

func addLead(lead Lead) (*Lead, error) {
	if err := data.CreateWithAutoIncrement(bucketLead, &lead); err != nil {
		return nil, err
	}
	return &lead, nil
}

func getLead(id uint64) (*Lead, error) {
	var l Lead
	if err := data.Get(bucketLead, data.IntToByteArray(id), &l); err != nil {
		return nil, err
	}
	return &l, nil
}

func updateLead(id uint64, l Lead) {
}

func removeLead(id uint64) error {
	return data.Delete(bucketLead, data.IntToByteArray(id))
}
