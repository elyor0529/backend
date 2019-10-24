package data

import (
	"testing"
)

type tmp struct {
	ID   uint64
	Name string
}

func (t *tmp) SetID(id uint64) {
	t.ID = id
}

func Test_DB_CreateWithAutoIncrement(t *testing.T) {
	x := &tmp{Name: "testing"}
	if err := CreateWithAutoIncrement(bucketTest, x); err != nil {
		t.Fatal(err)
	}

	var exists tmp
	if err := Get(bucketTest, IntToByteArray(x.ID), &exists); err != nil {
		t.Fatal(err)
	} else if x.Name != exists.Name {
		t.Errorf("expected name to be %s got %s", x.Name, exists.Name)
	}
}
