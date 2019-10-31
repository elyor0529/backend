package data

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var (
	ErrNotFound = fmt.Errorf("unable to find this key")
)

type AutoIncrementer interface {
	SetID(id uint64)
}

func CreateWithAutoIncrement(bucketName []byte, v AutoIncrementer) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		v.SetID(id)
		buf, err := Encode(v)
		if err != nil {
			return err
		}
		return b.Put(IntToByteArray(id), buf)
	})
	if err != nil {
		return err
	}
	return nil
}

func Get(bucketName, key []byte, v interface{}) error {
	var buf []byte
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		buf = b.Get(key)
		return nil
	})
	if err != nil {
		return err
	} else if len(buf) == 0 {
		return ErrNotFound
	}

	return Decode(buf, v)
}

// TODO: Can we use only the []interface and not both params.
func List(bucketName []byte, f func(buf []byte) error) error {
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		c := b.Cursor()

		for key, val := c.First(); key != nil; key, val = c.Next() {
			if err := f(val); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func Update(bucketName, key []byte, v interface{}) error {
	buf, err := Encode(v)
	if err != nil {
		return err
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put(key, buf)
	})
	return err
}

func Delete(bucketName, key []byte) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete(key)
	})
	return err
}
