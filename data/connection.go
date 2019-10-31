/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package data

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Open(dbFile string) error {
	if err := os.MkdirAll(path.Dir(dbFile), 0777); err != nil {
		log.Fatalf("unable to create database file %s err: %v", path.Dir(dbFile), err)
	}

	db, err := bolt.Open(dbFile, 0777, nil)
	if err != nil {
		return fmt.Errorf("unable to open database file %s: %v", dbFile, err)
	}

	DB = db

	return nil
}

func Close() {
	DB.Close()
}
