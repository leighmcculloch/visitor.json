package main


import (
  "fmt"
  "net/http"
  "encoding/json"
)

type MainResponse struct {
  AcceptedLanguages []string `json:"acceptedLanguages"`
  Readme string `json:"readme,omitempty"`
}

func MakeJson(r *http.Request, mr *MainResponse, pretty bool) (string, string) {
  var jsonBytes []byte
  if (pretty) {
    mr.Readme = "/readme"
    jsonBytes, _ = json.MarshalIndent(mr, "", "  ")
  } else {
    jsonBytes, _ = json.Marshal(mr)
  }
  return string(jsonBytes), "application/json"
}

func MakeJavascript(r *http.Request, mr *MainResponse, pretty bool) (string, string) {
  response, _ := MakeJson(r, mr, pretty)
  response = "window.visitor = " + response
  return response, "text/javascript"
}

func MakeJsonP(r *http.Request, mr *MainResponse, pretty bool) (string, string) {
  callback := r.URL.Query().Get("callback")
  if callback == "" {
    callback = "callback"
  }
  response, _ := MakeJson(r, mr, pretty)
  response = callback + "(" + response + ")"
  return response, "text/javascript"
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
  format := r.URL.Path[1:]
  mr := &MainResponse{}
  GetAcceptLanguageResponse(r, mr)

  var response string
  var contentType string

  switch format {
  case "jsonp":
    response, contentType = MakeJsonP(r, mr, false)
  case "js":
    response, contentType = MakeJavascript(r, mr, false)
  case "json":
    response, contentType = MakeJson(r, mr, false)
  case "":
    response, contentType = MakeJson(r, mr, true)
  }

  if response != "" {
    w.Header().Set("Content-Type", contentType)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Headers", "Content-type")
    fmt.Fprintln(w, response)
  } else {
    http.NotFound(w, r)
  }
}
