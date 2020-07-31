package export

import (
	"testing"

	model "github.com/wyllisMonteiro/go_mvc/models"
)

func TestNewContext(t *testing.T) {
	writer, request := CreateFakeHandler()
	var handlerServer HandlerServer = HandlerServer{
		Writer:  writer,
		Request: request,
	}

	var csv ExportArticles = CSV{}
	articles := model.InitArticles(1, "First title", "First description")
	articles = model.AddArticleItem(articles, 2, "Second title", "Second description")

	context1 := NewContext(csv, articles, handlerServer)

	t.Run("Articles not empty", func(t *testing.T) {
		if len(context1.Articles) <= 0 {
			t.Errorf("articles = %v, want 2", len(context1.Articles))
		}
	})

	var xlsx ExportArticles = XLSX{}
	var emptyArticles []model.Article = model.EmptyArticles()

	context2 := NewContext(xlsx, emptyArticles, handlerServer)

	t.Run("Articles empty", func(t *testing.T) {
		if len(context2.Articles) > 0 {
			t.Errorf("articles = %v, want 0", len(context2.Articles))
		}
	})

}
