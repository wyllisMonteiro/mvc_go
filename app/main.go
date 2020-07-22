package main

import (
	"github.com/wyllisMonteiro/go_mvc/app/router"
)

func main() {
	
	// Update - update article's price to 2000
	//db.Model(&article).Update("Title", "Deuxième")

	// Delete - delete article
	//db.Delete(&article)

	router.InitRoutes()
}