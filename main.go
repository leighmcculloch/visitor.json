package main

import (
  "log"
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/russross/blackfriday"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, r.URL.Path[1:])
}

func MarkdownHandler(w http.ResponseWriter, r *http.Request, filename string) {
  data, _ := ioutil.ReadFile(filename)
  html := string(blackfriday.MarkdownCommon(data))
  html = "<!doctype html><html><head><link href=markdown.css rel=stylesheet /></head><body class=markdown-body>" + html + "</body></html>"
  w.Header().Set("Content-Type", "text/html; utf-8")
  fmt.Fprintln(w, html)
}

func main() {
  // main api
  http.HandleFunc("/", MainHandler)

  // static pages
  http.HandleFunc("/favicon.ico", FileHandler)
  http.HandleFunc("/markdown.css", FileHandler)
  http.HandleFunc("/readme", func(w http.ResponseWriter, r *http.Request) {
    MarkdownHandler(w, r, "README.md")
  })

  var port string = os.Getenv("PORT")
  var address string
  if port == "" {
    port =  "5001"
    address = "127.0.0.1"
  }
  log.Println(fmt.Sprintf("Listening on %s:%s", address, port))
  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil))
}
