package controllers

import (
  //"os"
  "text/template"
  "fmt"
  "net/http"
  "log"
  model "github.com/wyllisMonteiro/go_mvc/app/models"
)

func GetWiki(w http.ResponseWriter, req *http.Request) {
  article := model.GetWiki(1)
  fmt.Println(article.Title)

  t := template.New("wikis")
  t = template.Must(t.ParseFiles("./templates/wikis.tmpl"))
  err := t.ExecuteTemplate(w, "wikis", nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }

  fmt.Println("test 1")
}