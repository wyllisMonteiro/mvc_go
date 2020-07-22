package main

import (
	"github.com/wyllisMonteiro/go_mvc/router"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Update - update article's price to 2000
	//db.Model(&article).Update("Title", "Deuxième")

	// Delete - delete article
	//db.Delete(&article)
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}