/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package chat

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

var (
	bucketActive = []byte("convo_active")
	bucketClosed = []byte("convo_closed")
)

type Conversation struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"userId"`
	UserName  string `json:"userName"`
	AgentID   uint64 `json:"agentId"`
	AgentName string `json:"agentName"`

	Created time.Time `json:"created"`
}

func NewConversation(c Conversation) (Conversation, error) {
	err := data.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketActive)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		c.ID = id
		buf, err := data.Encode(c)
		if err != nil {
			return err
		}
		return b.Put(data.IntToByteArray(id), buf)
	})
	return c, err
}
