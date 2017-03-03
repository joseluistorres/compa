package main

import (
  "encoding/json"
  "log"
  "net/http"
  "strconv"
  re "gopkg.in/gorethink/gorethink.v3"
  // "time"
  // "net/http/httputil"
  // "fmt"
  // "os"
)

type WebhookResponse struct {
  Username string `json:"username"`
  Text     string `json:"text"`
}

var (
  session *re.Session
)

// func debug(data []byte, err error) {
//     if err == nil {
//         fmt.Printf("%s\n\n", data)
//     } else {
//         log.Fatalf("%s\n\n", err)
//     }
// }

func init() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // data, err := httputil.DumpRequestOut(r, true)
    // log.Printf("%s\n\n", data)
    // if err != nil {
    //     log.Fatal(err)
    //   }
    incomingText := r.PostFormValue("text")
    if incomingText != "" && r.PostFormValue("user_id") != "" {
      text := parseText(incomingText)
      trigger_word := r.PostFormValue("trigger_word")
      log.Printf("Handling incoming request: %s", text)

      if trigger_word == "apuntale" {
        var response WebhookResponse
        response.Username = botUsername
        response.Text = "This is just a simulated response"
        log.Printf("Sending response: %s", response.Text)
        // testing the rethinkDB
        // share := NewShare("https://www.rethinkdb.com/docs/reql-data-exploration/", incomingText)
        // share.Created = time.Now()

        // // Insert the new item into the database
        // _, err := re.Table("shares").Insert(share).RunWrite(session)
        // if err != nil {
        //   http.Error(w, err.Error(), http.StatusInternalServerError)
        //   return
        // }

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
  log.Printf("Starting RethinkDB...")
  initDB()
  log.Printf("Starting HTTP server on %d", port)
  err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func initDB() {
  //var err error

  // session, err = re.Connect(re.ConnectOpts{
  //   Address: "127.0.0.1:28015",
  //   Database: os.Getenv("RETHINKDB_DATABASE"),
  //   Username: os.Getenv("RETHINKDB_USERNAME"),
  //   Password: os.Getenv("RETHINKDB_PASSWORD"),
  // })

  // if err != nil {
  //   log.Fatalln(err.Error())
  // }
}