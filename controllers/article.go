package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	model "github.com/wyllisMonteiro/go_mvc/models"
	service "github.com/wyllisMonteiro/go_mvc/services"
	export "github.com/wyllisMonteiro/go_mvc/services/export"
)

/**
 *
 * Render all articles view
 * GET /
 * GET /articles
 *
 */
func GetArticles(w http.ResponseWriter, req *http.Request) {
	service.RenderArticles(w)
}

/**
 *
 * Download articles to CSV or XLSX
 * POST /articles/download
 *
 */
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

/**
 *
 * Render an article view
 * GET /article/{id}
 *
 */
func GetArticle(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	article_id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderArticle(w, article_id)
}

/**
 *
 * Render create article view
 * GET /article/create
 *
 */
func CreateArticleForm(w http.ResponseWriter, req *http.Request) {
	service.RenderCreateArticle(w)
}

/**
 *
 * Create article and redirect to new article created
 * POST /article/create
 *
 */
func CreateArticle(w http.ResponseWriter, req *http.Request) {
	var article model.Article
	article.Title = req.FormValue("title")
	article.Description = req.FormValue("description")

	article_id, err := service.CreateArticle(article)
	if err != nil {
		log.Fatalf("Database : %s", err)
		return
	}

	service.Redirect(w, req, "/article/"+strconv.Itoa(article_id))
}

/**
 *
 * Render edit article view
 * GET /article/{id}/edit
 *
 */
func EditArticleForm(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	article_id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		log.Fatalf("Convert string to int: %s", err)
		return
	}

	service.RenderEditArticle(w, article_id)
}

/**
 *
 * Edit article and redirect to edited article
 * POST /article/{id}/edit
 *
 */
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

	service.Redirect(w, req, "/article/"+urlParams["id"])
}
