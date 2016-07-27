package main

import (
  "net/http"
  "github.com/russross/blackfriday"
  "os"
)

var port = os.Getenv("PORT")

func main() {
  if port == "" {
    port = "8080"
  }
  http.HandleFunc("/markdown", GenerateMarkdown)
  http.Handle("/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":" + port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request)  {
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}
