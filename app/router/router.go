package router

import (
	"log"
	"github.com/wyllisMonteiro/go_mvc/app/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetWikis).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}
