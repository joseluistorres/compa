package main

import (
  "time"
)

type Share struct {
  Id      string `gorethink:"id,omitempty"`
  Link    string
  Description  string
  Created time.Time
}

func NewShare(link string, description string) *Share {
  return &Share{
    Link:   link,
    Description: description,
  }
}

// type Issue struct {
//   Id      string `gorethink:"id,omitempty"`
//   ExternalId int
//   Subject    string `gorethink:"subject" json:"subject"`
//   Description string `gorethink:"description" json:"description"`
//   Status  string
//   Created time.Time
// }

// func NewIssue(external_id int, subject string, description string) *Issue {
//   return &Issue{
//     ExternalId: external_id,
//     Subject: subject,
//     Description: description,
//     Status: "active",
//   }
// }
