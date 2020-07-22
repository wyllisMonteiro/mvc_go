package controllers

import (
  //"os"
  "text/template"
  "fmt"
  "net/http"
  "log"
)

func GetWikis(w http.ResponseWriter, req *http.Request) {
  //var allFiles []string
  //allFiles = append(allFiles, "./templates/wikis.tmpl")

  //templates, err := template.ParseFiles(allFiles...)
  //if err != nil {
  //  fmt.Println(err.Error())
  //}

  //s1 := templates.Lookup("wikis.tmpl")
  //err = s1.ExecuteTemplate(w, "wikis", nil)
  //if err != nil {
  //  log.Fatalf("execution failed: %s", err)
  //}

  t := template.New("wikis")
  t = template.Must(t.ParseFiles("./templates/wikis.tmpl"))
  err := t.ExecuteTemplate(w, "wikis", nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }

  fmt.Println("test 1")
}