package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	model "github.com/wyllisMonteiro/go_mvc/models"
	export "github.com/wyllisMonteiro/go_mvc/services/export"
)

// Call database to get all articles
func GetArticles() ([]model.Article, error) {
	articles, err := model.GetArticles()
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return []model.Article{}, err
	}

	return articles, nil
}

// Call database to get article
func GetArticle(article_id int) (model.Article, error) {
	article, err := model.GetArticle(article_id)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return model.Article{}, err
	}

	return article, nil
}

// Call database to create article
func CreateArticle(article model.Article) (int, error) {
	article_id, err := model.CreateArticle(article)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return 0, err
	}

	return article_id, nil
}

// Call database to update articles
func UpdateArticle(article model.Article, new_article model.Article) error {
	err := model.EditArticle(article, new_article)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return err
	}

	return nil
}

// Display all articles page in a browser
func RenderArticles(w http.ResponseWriter) {
	articles, err := GetArticles()
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/articles.tmpl", "web/base.tmpl"},
		Writer:     w,
		Data:       articles,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Display an article page in a browser
func RenderArticle(w http.ResponseWriter, article_id int) {
	article, err := GetArticle(article_id)
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/article.tmpl", "web/base.tmpl"},
		Writer:     w,
		Data:       article,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Display create article page in a browser
func RenderCreateArticle(w http.ResponseWriter) {
	var render Render = Render{
		ParseFiles: []string{"web/create_article.tmpl", "web/base.tmpl"},
		Writer:     w,
		Data:       nil,
	}

	err := RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Display edit article page in a browser
func RenderEditArticle(w http.ResponseWriter, article_id int) {
	article, err := GetArticle(article_id)
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/edit_article.tmpl", "web/base.tmpl"},
		Writer:     w,
		Data:       article,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Export articles to CSV or XLSX
func LaunchExport(type_export string, handler export.HandlerServer) error {
	var export_file *export.Context

	articles, err := GetArticles()
	if err != nil {
		return err
	}

	switch type_export {
	case "csv":
		csv := &export.CSV{}
		export_file = export.NewContext(csv, articles, handler)
	case "xlsx":
		xlsx := &export.XLSX{}
		export_file = export.NewContext(xlsx, articles, handler)
	default:
		return errors.New("Unavaliable file format")
	}

	err = export_file.MakeExport()
	if err != nil {
		return err
	}

	return nil
}
