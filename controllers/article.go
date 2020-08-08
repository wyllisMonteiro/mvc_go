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
	typeDownload := req.FormValue("type_download")

	handler := export.HandlerServer{
		Writer:  w,
		Request: req,
	}

	err := service.LaunchExport(typeDownload, handler)
	/*if err != nil {
		return
	}*/
}

// Render an article view
func GetArticle(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	articleId, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderArticle(w, articleId)
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

	articleId, err := service.CreateArticle(article)
	if err != nil {
		log.Fatalf("Database : %s", err)
		return
	}

	route := fmt.Sprintf("/article/%d", articleId)
	service.Redirect(w, req, route)
}

// Render edit article view
func EditArticleForm(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	articleId, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderEditArticle(w, articleId)
}

// Edit article and redirect to edited article
func UpdateArticle(w http.ResponseWriter, req *http.Request) {
	var newArticle model.Article
	newArticle.Title = req.FormValue("title")
	newArticle.Description = req.FormValue("description")

	urlParams := mux.Vars(req)
	articleId, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	article, err := service.GetArticle(articleId)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return
	}

	err = service.UpdateArticle(article, newArticle)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
	}

	route := fmt.Sprintf("/article/%s", urlParams["id"])
	service.Redirect(w, req, route)
}
