package controllers

import (
  "text/template"
  "fmt"
  "net/http"
  "log"
  "strconv"
	"github.com/gorilla/mux"
  model "github.com/wyllisMonteiro/go_mvc/models"
)

type Page struct {
  Content  interface{}
}

func GetArticles(w http.ResponseWriter, req *http.Request) {
  articles, err := model.GetArticles()

  if err != nil {
    log.Fatalf("Model execution: %s", err)
    return
  }

  page := Page{articles}

  tmpl, err := template.ParseFiles("web/articles.tmpl", "web/base.tmpl")
  err = tmpl.Execute(w, page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}

func GetArticle(w http.ResponseWriter, req *http.Request) {
  //article_id, err := strconv.Atoi(req.FormValue("id"))
  urlParams := mux.Vars(req)
  article_id, err := strconv.Atoi(urlParams["id"])
  if err != nil {
    log.Fatalf("Convert string to int: %s", err)
    return
  }

  article, err := model.GetArticle(article_id)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    return
  }

  page := Page{article}
  fmt.Println(article.Title)

  t, err := template.ParseFiles("./web/article.tmpl", "./web/base.tmpl",)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }

  err = t.Execute(w, page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}