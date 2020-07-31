package models

import (
	"fmt"
)

// Structure of an article in db
type Article struct {
	ID          int
	Title       string
	Description string
}

// Get all articles or error
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

// Get one article by id or error
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

// Create an article or error
func CreateArticle(article Article) (int, error) {
	db, err := ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	defer db.Close()

	db.AutoMigrate(&Article{})

	db.Save(&article)

	return article.ID, nil
}

// Edit an article or error
func EditArticle(article Article, newArticle Article) error {
	db, err := ConnectToBDD()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer db.Close()

	db.AutoMigrate(&Article{})

	db.Model(&article).Updates(newArticle)

	return nil
}
