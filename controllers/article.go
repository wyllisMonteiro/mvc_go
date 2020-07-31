package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	model "github.com/wyllisMonteiro/go_mvc/models"
	service "github.com/wyllisMonteiro/go_mvc/services"
	export "github.com/wyllisMonteiro/go_mvc/services/export"
)

// Render all articles view
func GetArticles(w http.ResponseWriter, req *http.Request) {
	service.RenderArticles(w)
}

// Download articles to CSV or XLSX
func DownloadArticles(w http.ResponseWriter, req *http.Request) {
	type_download := req.FormValue("type_download")

	handler := export.HandlerServer{
		Writer:  w,
		Request: req,
	}

	err := service.LaunchExport(type_download, handler)
	if err != nil {
		return
	}
}

// Render an article view
func GetArticle(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	article_id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderArticle(w, article_id)
}

// Render create article view
func CreateArticleForm(w http.ResponseWriter, req *http.Request) {
	service.RenderCreateArticle(w)
}

// Create article and redirect to new article created
func CreateArticle(w http.ResponseWriter, req *http.Request) {
	var article model.Article
	article.Title = req.FormValue("title")
	article.Description = req.FormValue("description")

	article_id, err := service.CreateArticle(article)
	if err != nil {
		log.Fatalf("Database : %s", err)
		return
	}

	route := fmt.Sprintf("/article/%d", article_id)
	service.Redirect(w, req, route)
}

// Render edit article view
func EditArticleForm(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	article_id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderEditArticle(w, article_id)
}

// Edit article and redirect to edited article
func UpdateArticle(w http.ResponseWriter, req *http.Request) {
	var new_article model.Article
	new_article.Title = req.FormValue("title")
	new_article.Description = req.FormValue("description")

	urlParams := mux.Vars(req)
	article_id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	article, err := service.GetArticle(article_id)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return
	}

	err = service.UpdateArticle(article, new_article)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
	}

	route := fmt.Sprintf("/article/%s", urlParams["id"])
	service.Redirect(w, req, route)
}
