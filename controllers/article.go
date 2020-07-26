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

type PageEditArticle struct {
  Status string
  Message string
  Data interface{}
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

func CreateArticleForm(w http.ResponseWriter, req *http.Request) {
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

func EditArticleForm(w http.ResponseWriter, req *http.Request) {
  page_article := PageEditArticle{}

  urlParams := mux.Vars(req)
  article_id, err := strconv.Atoi(urlParams["id"])
  if err != nil {
    log.Fatalf("Convert string to int: %s", err)
    page_article = PageEditArticle{"", "", nil}

  }

  article, err := model.GetArticle(article_id)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    page_article = PageEditArticle{"", "", nil}
  }

  page_article.Status = ""
  page_article.Message = ""
  page_article.Data = article

  tmpl, err := template.ParseFiles("web/edit_article.tmpl", "web/base.tmpl")

  err = tmpl.Execute(w, page_article)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

func EditArticle(w http.ResponseWriter, req *http.Request) {
  var new_article model.Article
  new_article.Title = req.FormValue("title")
  new_article.Description = req.FormValue("description")

  page_article := PageEditArticle{"success", "L'article à bien été modifié", nil}

  urlParams := mux.Vars(req)
  article_id, err := strconv.Atoi(urlParams["id"])
  if err != nil {
    log.Fatalf("Convert string to int: %s", err)
    page_article = PageEditArticle{"danger", "Une erreur est survenue avec la récupération des paramètres de l'url", nil}
  }

  article, err := model.GetArticle(article_id)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    page_article = PageEditArticle{"danger", "Impossible de récupérer les informations de l'article", nil}
  }

  err = model.EditArticle(article, new_article)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    page_article = PageEditArticle{"danger", "Impossible de récupérer les informations de l'article", nil}
  }

  tmpl, err := template.ParseFiles("web/edit_article.tmpl", "web/base.tmpl")

  err = tmpl.Execute(w, page_article)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

}