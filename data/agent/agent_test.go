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

	"github.com/parle-io/backend/data"
)

func TestMain(m *testing.M) {
	if err := data.Open("../db/test.db"); err != nil {
		log.Fatal(err)
	}

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

	if err := Add(agent); err != nil {
		t.Fatal(err)
	}

	check, err := GetByEmail(agent.Email)
	if err != nil {
		t.Fatal(err)
	} else if check.Email != agent.Email {
		t.Errorf("expected email %s got %s", agent.Email, check.Email)
	}

}
