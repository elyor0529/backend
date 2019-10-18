/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package agent

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
)

func TestMain(m *testing.M) {
	if err := data.Open("../db/test.db"); err != nil {
		log.Fatal(err)
	}

	data.DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket(bucketAgents); err != nil {
			log.Fatal(err)
		}
		return nil
	})

	retval := m.Run()
	data.Close()
	os.Exit(retval)
}

func Test_Agent_Add(t *testing.T) {
	agent := Agent{
		Email:    "unit@test.com",
		Password: "unit-test",
		Created:  time.Now(),
	}

	if _, err := Add(agent); err != nil {
		t.Fatal(err)
	}

	check, err := GetByEmail(agent.Email)
	if err != nil {
		t.Fatal(err)
	} else if check.Email != agent.Email {
		t.Errorf("expected email %s got %s", agent.Email, check.Email)
	}

}
