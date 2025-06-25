package main

import (
	"encoding/json",
	"log",
	"net/http",
	"string",
	"time",
	"context",
	"os",
	"os/signal",
	"syscall",

	"github.com/go-chi/chi",
	"github.com/go-chi/chi/middleware",
	"github.com/go-chi/cors",
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2",
	"gopkg.in/mgo.v2/bson",
)

var rnd *renderer.Render
var db *mgo.Database

const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
)

func init() {
	rnd = renderer.New()
	srv := os.Getenv(DB_HOST) + ":" + os.Getenv(DB_PORT)
	session, err := mgo.Dial(srv)
	if err != nil {
		panic(err)
	}
	db = session.DB(os.Getenv(DB_NAME))
}