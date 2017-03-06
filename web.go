package main

import (
  "encoding/json"
  "log"
  "net/http"
  "strconv"
  re "gopkg.in/gorethink/gorethink.v3"
  "time"
  "os"
)

type WebhookResponse struct {
  Username string `json:"username"`
  Text     string `json:"text"`
}

var (
  session *re.Session
)

func init() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if session == nil {
      initDB()
    }
    incomingText := r.PostFormValue("text")
    log.Printf("RAW TEXT: %s", incomingText)
    if incomingText != "" && r.PostFormValue("user_id") != "" {
      text := parseText(incomingText)
      url_match := getUrl(incomingText)
      user_slack_id := getUserSlackId(incomingText)
      trigger_word := r.PostFormValue("trigger_word")
      log.Printf("user id: %s", r.PostFormValue("user_id"))
      log.Printf("user name: %s", r.PostFormValue("user_name"))
      log.Printf("token id: %s", r.PostFormValue("token"))
      log.Printf("Handling incoming request: %s", text)

      if trigger_word == "apuntale" && r.PostFormValue("user_name") == "joseluistorres" {
        var response WebhookResponse
        response.Username = botUsername
        response.Text = "That comment/link has been saved"
        log.Printf("Sending response: %s", response.Text)
        log.Println(url_match)
        log.Println(user_slack_id)
        /*
         * Create or get User
         */
         user := NewUser("joseluistorres", "@gdljs", "", user_slack_id)
         //user.Created = time.Now()

         resp, err := re.Table("users").Insert(user, re.InsertOpts{
              Conflict: "replace",
           }).RunWrite(session)

         if err != nil {
          log.Println("----------error 1--------------")
          log.Println(err.Error())
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
         }

        log.Println(resp.GeneratedKeys)

        /*
         * Create a new share link
         *
         */
        share := NewShare(resp.GeneratedKeys[0], url_match, "This is just a generic description")
        share.Created = time.Now()

        // Insert the new item into the database
        _, err = re.Table("shares").Insert(share).RunWrite(session)
        if err != nil {
          log.Fatal(err)
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }

        //resp.GeneratedKeys

        b, err := json.Marshal(response)
        if err != nil {
          log.Fatal(err)
        }

        w.Write(b)
      }
    }
  })
}

func StartServer(port int) {
  log.Printf("Starting HTTP server on %d", port)
  err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

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