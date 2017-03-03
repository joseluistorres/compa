package main

import (
  "time"
)

type Share struct {
  Id      string `gorethink:"id,omitempty"`
  UserId  string `gorethink:"user_id" json:"user_id"`
  Link    string `gorethink:"link" json:"link"`
  Description  string `gorethink:"description" json:"description"`
  Created time.Time
}

func NewShare(user_id string, link string, description string) *Share {
  return &Share{
    UserId: user_id,
    Link:   link,
    Description: description,
  }
}

type User struct {
  Id       string `gorethink:"id,omitempty"`
  SlackUsername string `gorethink:"username" json:"username"`
  Context  string `gorethink:"context" json:"context"`
  Twitter   string `gorethink:"twitter" json:"twitter"`
  Slackid   string `gorethink:"slackid" json:"slackid"`
  // Created time.Time
}

func NewUser(slack_user_name string, context string, twitter string, slack_id string) *User {
  return &User{
    SlackUsername: slack_user_name,
    Context: context,
    Twitter: twitter,
    Slackid: slack_id,
  }
}
