/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */
package data

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestMain(m *testing.M) {
	os.Remove("db/test.db")

	retval := m.Run()

	Close()
	os.Exit(retval)
}

func Test_Data_Connection(t *testing.T) {
	if err := Open("db/test.db"); err != nil {
		t.Fatal(err)
	}

	// Add all needed buckets
	DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("agents"))
		return err
	})
}
