/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/parle-io/backend/data"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	db   = flag.String("db", "db/parle.db", "database path and file name (default to db/parle.db)")
)

func main() {
	flag.Parse()

	// Initiating the database connection pool
	if err := data.Open(*db); err != nil {
		log.Fatal(err)
	}

	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("unable to start server: ", err)
	}
}
