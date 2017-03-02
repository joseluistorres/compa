package main

import (
  "encoding/json"
  "log"
  "net/http"
  "strconv"
)

type WebhookResponse struct {
  Username string `json:"username"`
  Text     string `json:"text"`
}

func init() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
