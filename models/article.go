package models

import (
	"fmt"
	//"github.com/jinzhu/gorm"
)

type Article struct {
	ID int
	Title string
	Description string
}

func GetArticles() ([]Article, error) {
	var articles []Article

	db, err := ConnectToBDD()
	if err != nil {
		fmt.Println(err.Error())
		return articles, err
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	err = db.Find(&articles).Error // find articles
	if err != nil {
		fmt.Println(err.Error())
		return articles, err
	}

	fmt.Println(articles[2].ID)

	return articles, nil
}


func GetArticle(id int) (Article, error) {
	var article Article

	db, err := ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return article, err
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.First(&article, id) // find article with id

	return article, nil
}

func CreateArticle() {
	db, err := ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.Create(&Article{Title: "Premier article", Description: "test 2"})
}