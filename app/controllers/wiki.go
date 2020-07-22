package controllers

import (
  "text/template"
  "fmt"
  "net/http"
  "log"
  model "github.com/wyllisMonteiro/go_mvc/app/models"
)

type PageWikis struct {
  Title    string
  Articles []model.Article
}

func GetWikis(w http.ResponseWriter, req *http.Request) {
  articles := model.GetWikis()

  page := PageWikis{"Titre de ma page", articles}

  tmpl := template.New("wikis")
  tmpl = template.Must(tmpl.ParseFiles("./templates/wikis.tmpl", "templates/articles.tmpl"))
  err := tmpl.ExecuteTemplate(w, "wikis", page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}

func GetWiki(w http.ResponseWriter, req *http.Request) {
  article := model.GetWiki(1)
  fmt.Println(article.Title)

  t := template.New("wikis")
  t = template.Must(t.ParseFiles("./templates/wikis.tmpl"))
  err := t.ExecuteTemplate(w, "wikis", nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}