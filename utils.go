package main

import (
  "regexp"
  "strings"
  "github.com/mvdan/xurls"
)

var (
  messageRegex *regexp.Regexp
)

func init() {
  messageRegex = regexp.MustCompile(`<([^>]+)>`)
}

func parseText(text string) string {
  matches := messageRegex.FindAllStringSubmatch(text, -1)
  for _, matches2 := range matches {

    if strings.HasPrefix(matches2[1], "http") || strings.HasPrefix(matches2[1], "mailto") {
      text = strings.Replace(text, matches2[0], "", -1)

    } else if strings.HasPrefix(matches2[1], "@U") {
      parts := strings.SplitN(matches2[1], "|", 2)

      if len(parts) == 2 {
        text = strings.Replace(text, matches2[0], "@"+parts[1], -1)
      } else {
        text = strings.Replace(text, matches2[0], "", -1)
      }

    } else if strings.HasPrefix(matches2[1], "@") {
      text = strings.Replace(text, matches2[0], matches2[1], -1)

    } else if strings.HasPrefix(matches2[1], "#") {
      parts := strings.SplitN(matches2[1], "|", 2)

      if len(parts) == 2 {
        text = strings.Replace(text, matches2[0], "#"+parts[1], -1)
      } else {
        text = strings.Replace(text, matches2[0], "", -1)
      }

    }
  }

  text = strings.TrimSpace(text)

  text = strings.Replace(text, "&lt;", "<", -1)
  text = strings.Replace(text, "&gt;", ">", -1)
  text = strings.Replace(text, "&amp;", "&", -1)

  return text
}

func getUrl(text string) string {
  return xurls.Relaxed.FindString(text)
}

func getUserSlackId(text string) string {
  matches := messageRegex.FindAllStringSubmatch(text, -1)
  get_user := ""
  for _, sub_str := range matches {
    if strings.HasPrefix(sub_str[1], "@U") {
      get_user = sub_str[1]
    }
  }
  return get_user
}