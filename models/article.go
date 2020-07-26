package models

import (
	"fmt"
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

func CreateArticle(article Article) (error) {
	db, err := ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.Create(&article)
	
	return nil
}