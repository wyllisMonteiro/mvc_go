package controllers

import (
  "text/template"
  "net/http"
  "log"
  "strconv"
	"github.com/gorilla/mux"
  model "github.com/wyllisMonteiro/go_mvc/models"
)

type PageShowArticle struct {
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

  send_data := PageShowArticle{articles}

  tmpl, err := template.ParseFiles("web/articles.tmpl", "web/base.tmpl")
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, send_data)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
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

  send_data := PageShowArticle{article}

  tmpl, err := template.ParseFiles("./web/article.tmpl", "./web/base.tmpl",)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, send_data)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

/**
 *
 * Display article form before posting data
 *
 */
func CreateArticleForm(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("web/create_article.tmpl", "web/base.tmpl")
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, nil)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

/**
 *
 * Data posted
 *
 */
func CreateArticle(w http.ResponseWriter, req *http.Request) {
  var article model.Article
  article.Title = req.FormValue("title")
  article.Description = req.FormValue("description")

  send_data := PageCreateArticle{"success", "L'article à bien été créé"}

  err := model.CreateArticle(article)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    send_data = PageCreateArticle{"error", "Une erreur est survenue lors de la création de l'article"}
  }
  
  tmpl, err := template.ParseFiles("web/create_article.tmpl", "web/base.tmpl")
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, send_data)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

/**
 *
 * Display article form before posting data
 *
 */
func EditArticleForm(w http.ResponseWriter, req *http.Request) {
  send_data := PageEditArticle{}

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

  send_data = PageEditArticle{"", "", article}

  tmpl, err := template.ParseFiles("web/edit_article.tmpl", "web/base.tmpl")
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, send_data)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }
}

/**
 *
 * Data posted
 *
 */
func EditArticle(w http.ResponseWriter, req *http.Request) {
  var new_article model.Article
  new_article.Title = req.FormValue("title")
  new_article.Description = req.FormValue("description")

  send_data := PageEditArticle{"success", "L'article à bien été modifié", nil}

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

  err = model.EditArticle(article, new_article)
  if err != nil {
    log.Fatalf("Model execution: %s", err)
    send_data = PageEditArticle{"danger", "Impossible de modifier les informations de l'article", nil}
  }

  tmpl, err := template.ParseFiles("web/edit_article.tmpl", "web/base.tmpl")
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

  err = tmpl.Execute(w, send_data)
  if err != nil {
    log.Fatalf("Template execution: %s", err)
    return
  }

}