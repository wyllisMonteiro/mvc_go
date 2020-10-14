package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/warticle/router"
)

/**
 *
 * Author : Wyllis Monteiro
 * Student at Hetic
 * Project MT4
 *
 * Learn software architecture
 *
 * Routes are in /router/router.go file
 * Each route load a controller (handler)
 *
 * Please check README.md
 *
 */
func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
