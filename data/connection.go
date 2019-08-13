/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package data

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Open(dbFile string) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return fmt.Errorf("unable to open database file %s: %v", dbFile, err)
	}

	DB = db

	return nil
}

func Close() {
	DB.Close()
}
