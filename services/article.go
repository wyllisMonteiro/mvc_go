package services

import (
	"log"
	"net/http"
	"fmt"
	export "github.com/wyllisMonteiro/go_mvc/services/export"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

func GetArticles() ([]model.Article, error) {
	articles, err := model.GetArticles()
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return []model.Article{}, err
	}

	return articles, nil
}

func GetArticle(article_id int) (model.Article, error) {
	article, err := model.GetArticle(article_id)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
	  	return model.Article{}, err
	}

	return article, nil
}

func CreateArticle(article model.Article) (int, error) {
	article_id, err := model.CreateArticle(article)
  	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return 0, err
	}

	return article_id, nil
}

func UpdateArticle(article model.Article, new_article model.Article) (error) {
	err := model.EditArticle(article, new_article)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return err
	}

	return nil
}

func RenderArticles(w http.ResponseWriter) {
	articles, err := GetArticles()
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/articles.tmpl", "web/base.tmpl"},
		Writter: w,
		Data: articles,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func RenderArticle(w http.ResponseWriter, article_id int) {
	article, err := GetArticle(article_id)
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/article.tmpl", "web/base.tmpl"},
		Writter: w,
		Data: article,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func RenderCreateArticle(w http.ResponseWriter) {
	var render Render = Render{
		ParseFiles: []string{"web/create_article.tmpl", "web/base.tmpl"},
		Writter: w,
		Data: nil,
	}

	err := RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func RenderEditArticle(w http.ResponseWriter, article_id int) {
	article, err := GetArticle(article_id)
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/edit_article.tmpl", "web/base.tmpl"},
		Writter: w,
		Data: article,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func LaunchExport(type_export string) {
	var export_file *export.Context

	articles, err := GetArticles()
	if err != nil {
		return
	} 

	switch type_export {
	case "csv":
		csv := &export.CSV{}
		export_file = export.NewContext(csv, articles)
	case "xlsx":
		xlsx := &export.XLSX{}
		export_file = export.NewContext(xlsx, articles)
	default:
		return
	}

	export_file.MakeExport()
}