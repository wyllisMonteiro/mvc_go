package router

import (
	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/go_mvc/controllers"
)

func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/", controllers.GetArticles).Methods("GET")
	r.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	r.HandleFunc("/articles/download", controllers.DownloadArticles).Methods("POST")
	r.HandleFunc("/article/create", controllers.CreateArticleForm).Methods("GET")
	r.HandleFunc("/article/create", controllers.CreateArticle).Methods("POST")
	r.HandleFunc("/article/{id}", controllers.GetArticle).Methods("GET")
	r.HandleFunc("/article/{id}/edit", controllers.EditArticleForm).Methods("GET")
	r.HandleFunc("/article/{id}/edit", controllers.UpdateArticle).Methods("POST")
	return r
}
