package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	model "github.com/wyllisMonteiro/warticle/models"
	export "github.com/wyllisMonteiro/warticle/services/export"
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
func GetArticle(articleId int) (model.Article, error) {
	article, err := model.GetArticle(articleId)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return model.Article{}, err
	}

	return article, nil
}

// Call database to create article
func CreateArticle(article model.Article) (int, error) {
	articleId, err := model.CreateArticle(article)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return 0, err
	}

	return articleId, nil
}

// Call database to update articles
func UpdateArticle(article model.Article, newArticle model.Article) error {
	err := model.EditArticle(article, newArticle)
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
func RenderArticle(w http.ResponseWriter, articleId int) {
	article, err := GetArticle(articleId)
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
func RenderEditArticle(w http.ResponseWriter, articleId int) {
	article, err := GetArticle(articleId)
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
func LaunchExport(typeExport string, handler export.HandlerServer) error {
	var exportFile *export.Context

	articles, err := GetArticles()
	if err != nil {
		return err
	}

	var file export.ExportArticles
	switch typeExport {
	case "csv":
		file = &export.CSV{}
	case "xlsx":
		file = &export.XLSX{}
	default:
		return errors.New("Unavaliable file format")
	}

	exportFile = export.NewContext(file, articles, handler)

	err = exportFile.MakeExport()
	if err != nil {
		return err
	}

	return nil
}
