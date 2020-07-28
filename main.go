package main

import (
	//"github.com/wyllisMonteiro/go_mvc/router"
	"log"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
	export "github.com/wyllisMonteiro/go_mvc/services/export"
)

func main() {
	_, _ = export.GetExport("csv")
	//fmt.Printf("Gun: %s", ak47.GetName())
	//r := mux.NewRouter()
	//router.InitRoutes(r)
	//log.Fatal(http.ListenAndServe(":9000", r))
	log.Fatal(http.ListenAndServe(":9000", nil))
}