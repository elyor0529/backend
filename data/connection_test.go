/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */
package data

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

var bucketTest = []byte("tests")

func TestMain(m *testing.M) {
	if err := os.Remove("db/test.db"); err != nil {
		fmt.Println("error reminving the test db", err)
	}

	if err := Open("db/test.db"); err != nil {
		log.Fatal(err)
	}

	err := DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket(bucketTest); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	retval := m.Run()
	Close()
	os.Exit(retval)
}
