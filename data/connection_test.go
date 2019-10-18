/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */
package data

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("test started")

	if err := os.Remove("db/test.db"); err != nil {
		fmt.Println("error reminving the test db", err)
	}

	fmt.Println("opening database")
	Open("db/test.db")

	fmt.Println("running the test")
	retval := m.Run()

	fmt.Println("closing the database")
	Close()
	os.Exit(retval)
}
