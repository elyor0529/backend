# Parle backend [![Build Status](https://travis-ci.org/parle-io/backend.svg?branch=master)](https://travis-ci.org/parle-io/backend) [![Code Climate](https://codeclimate.com/github/parle-io/backend/badges/gpa.svg)](https://codeclimate.com/github/parle-io/backend) [![Go Report Card](https://goreportcard.com/badge/github.com/parle-io/backend)](https://goreportcard.com/report/github.com/parle-io/backend)

# Backend

backend service that handles websocket and API for Parle (frontend as well as the widget).

## Build

This is a standard `Go` project. Dependencies are not vendored yet and no decision regarding the usage of Go's 1.11 
module has been taken.

You could do the following inside your `GOPATH`:

```shell
$ mkdir backend && cd backend
$ git clone https://github.com/parle-io/backend
$ go get
$ go build
```

## Usage

At this moment the program is mostly used with `go test` and not yet in runtime.

```shell
$ ./backend -addr ":8080"
```

**Parameters:**

`-addr` -> the address and port for the web server

When started the websocket endpoint is `addr`/ws.

