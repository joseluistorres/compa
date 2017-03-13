package main

import (
  "log"
  "os"
  "html/template"
  re "gopkg.in/gorethink/gorethink.v3"
)

func writeToHTML(){
  f, err := os.Create("index.html")
  if err != nil {
      panic(err)
  }

  defer f.Close()

  const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>TMIGDLJS</title>
  <link href='https://fonts.googleapis.com/css?family=Lato:400,900' rel='stylesheet' type='text/css'>
  <link rel="stylesheet" href="style.css">
</head>


   <body>

   <section id="main">
        <ul id="todo-list">
            {{range .}}
            <li data-id="{{.Id}}">
                <div class="view">
                    <a href="{{.Link}}" class="button toggle"></a>
                    <span>{{.Description}}</span>
                    <a href="/delete/{{.Id}}" class="button destroy"></a>
                </div>
            </li>
            {{end}}
        </ul>
    </section>

  </body>
</html>`


  t, err := template.New("webpage").Parse(tpl)
  check(err)

  var shares ShareList

  res, err := re.Table("shares").Run(session)
  err = res.All(&shares)

  check(err)

  defer res.Close()


  err = t.Execute(f, shares)
  check(err)

}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
