/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package data

import (
	"testing"
)

type dummy struct {
	OK   bool
	Test string
}

func Test_Data_Encoding(t *testing.T) {
	d := dummy{true, "unit"}

	b, err := Encode(d)
	if err != nil {
		t.Fatal(err)
	}

	var check dummy
	if err := Decode(b, &check); err != nil {
		t.Fatal(err)
	} else if check.OK == false {
		t.Errorf("expected OK to be true")
	} else if check.Test != "unit" {
		t.Errorf("expected Test to be unit got %s", check.Test)
	}
}
