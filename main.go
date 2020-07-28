package main

import (
	"log"
	"net/http"
	"github.com/wyllisMonteiro/go_mvc/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}