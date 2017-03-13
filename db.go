package main

import (
  "log"
  "os"
  re "gopkg.in/gorethink/gorethink.v3"
)

func initDB() {
  log.Printf("Starting RethinkDB connection from go...")
  var err error

  session, err = re.Connect(re.ConnectOpts{
    Address: os.Getenv("RETHINKDB_URL_PORT"),
    Database: os.Getenv("RETHINKDB_DATABASE"),
    Username: os.Getenv("RETHINKDB_USERNAME"),
    Password: os.Getenv("RETHINKDB_PASSWORD"),
  })

  if err != nil {
    log.Fatalln(err.Error())
  }
}