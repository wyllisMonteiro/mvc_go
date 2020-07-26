package controllers

import (
  "text/template"
  "net/http"
  "log"
  "strconv"
	"github.com/gorilla/mux"
  model "github.com/wyllisMonteiro/go_mvc/models"
)

type Page struct {
  Content interface{}
}

type PageCreateArticle struct {
  Status string
  Message string
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

  t, err := template.ParseFiles("./web/article.tmpl", "./web/base.tmpl",)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }

  err = t.Execute(w, page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
  }
}

func GetArticleForm(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("web/create_article.tmpl", "web/base.tmpl")

  err = tmpl.Execute(w, nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

func CreateArticle(w http.ResponseWriter, req *http.Request) {
  var article model.Article
  article.Title = req.FormValue("title")
  article.Description = req.FormValue("description")

  page_article := PageCreateArticle{"success", "L'article à bien été créé"}

  err := model.CreateArticle(article)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    page_article = PageCreateArticle{"error", "Une erreur est survenue lors de la création de l'article"}
  }
  

  page := Page{page_article}

  tmpl, err := template.ParseFiles("web/create_article.tmpl", "web/base.tmpl")

  err = tmpl.Execute(w, page)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}