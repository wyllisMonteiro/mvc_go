package main

import (
	"github.com/wyllisMonteiro/go_mvc/router"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Delete - delete article
	//db.Delete(&article)
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}