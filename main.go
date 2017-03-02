package main

// Main entry point for the app. Handles command-line options, starts the web

import (
  "flag"
  "fmt"
  // "log"
  "math/rand"
  "os"
  "time"
  "strconv"
)

var (
  httpPort       int
  botUsername    string
)

func init() {
  rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
}

func main() {
  // Parse command-line options
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "usage: ./compa -port=8000\n")
    flag.PrintDefaults()
  }
  port, err := strconv.Atoi(os.Getenv("PORT"))

  if err != nil {
    fmt.Errorf("Must have an ENV PORT VARIABLE RUNNING")
  }

  flag.IntVar(&httpPort, "port", port, "The HTTP port on which to listen")
  flag.StringVar(&botUsername, "botUsername", "patron", "The name of the bot when it speaks")

  flag.Parse()

  if httpPort == 0 {
    flag.Usage()
    os.Exit(2)
  }


  // Start the webserver
  StartServer(httpPort)
}
