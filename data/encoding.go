/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package data

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

func Encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte, v interface{}) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

// IntToByteArray returns an 8-byte big endian representation of v.
func IntToByteArray(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
