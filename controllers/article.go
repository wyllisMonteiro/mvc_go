package controllers

import (
  "text/template"
  "fmt"
  "net/http"
  "log"
  model "github.com/wyllisMonteiro/go_mvc/models"
)

type PageWikis struct {
  Title    string
  Articles []model.Article
}

func GetArticles(w http.ResponseWriter, req *http.Request) {
  articles, err := model.GetArticles()
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    return
  }

  page := PageWikis{"Titre de ma page", articles}

  tmpl := template.New("wikis")
  tmpl = template.Must(tmpl.ParseFiles("./web/wikis.tmpl", "web/articles.tmpl"))
  err = tmpl.ExecuteTemplate(w, "wikis", page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}

func GetArticle(w http.ResponseWriter, req *http.Request) {
  article := model.GetArticle(1)
  fmt.Println(article.Title)

  t := template.New("wikis")
  t = template.Must(t.ParseFiles("./web/wikis.tmpl"))
  err := t.ExecuteTemplate(w, "wikis", nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}