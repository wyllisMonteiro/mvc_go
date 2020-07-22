package router

import (
	"github.com/wyllisMonteiro/go_mvc/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) (*mux.Router) {
	r.HandleFunc("/", controllers.GetArticles).Methods("GET")
	//r.HandleFunc("/", controllers.GetArticle).Methods("GET")
	return r
}
