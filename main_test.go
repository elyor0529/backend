/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
	"github.com/parle-io/backend/data/agent"
)

type wstestdata struct {
	svr   *httptest.Server
	hub   *Hub
	url   string
	admin agent.Agent
}

var ws *wstestdata

func TestMain(m *testing.M) {
	// we make sure to delete the test db before starting
	os.Remove("db/test.db")

	if err := data.Open("db/test.db"); err != nil {
		log.Fatal(err)
	}

	// we creates the buckets needed for the tests
	data.DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket([]byte("agents")); err != nil {
			return err
		}
		if _, err := tx.CreateBucket([]byte("identity")); err != nil {
			return err
		}
		if _, err := tx.CreateBucket([]byte("visitors")); err != nil {
			return err
		}
		if _, err := tx.CreateBucket([]byte("leads")); err != nil {
			return err
		}
		if _, err := tx.CreateBucket([]byte("users")); err != nil {
			return err
		}
		return nil
	})

	// we create an agent to use during tests
	admin, err := agent.Add(agent.Agent{Email: "unit@test.com",
		Password: "unittest",
		Token:    "tokentest",
		Created:  time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}

	// we create a user to use during tests

	hub := newHub()
	go hub.run()

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}))
	defer svr.Close()

	ws = &wstestdata{
		svr:   svr,
		hub:   hub,
		url:   strings.Replace(svr.URL, "http", "ws", -1),
		admin: *admin,
	}

	ret := m.Run()
	os.Exit(ret)
}
