package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	config "github.com/wyllisMonteiro/go_mvc/app/config"
)

type Article struct {
	gorm.Model
	ID int
	Title string
	Description string
}

func GetWikis() ([]Article) {
	var articles []Article

	db, err := config.ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return articles
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	err = db.Find(&articles).Error // find articles
	if err != nil {
		fmt.Println(err.Error())
		return articles
	}

	return articles
}


func GetWiki(id int) (Article) {
	var article Article

	db, err := config.ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return article
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.First(&article, id) // find article with id

	return article
}

func CreateWiki() {
	db, err := config.ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.Create(&Article{Title: "Premier article", Description: "test 2"})
}