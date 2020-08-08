package export

import (
	"net/http/httptest"
	"testing"

	model "github.com/wyllisMonteiro/go_mvc/models"
)

// Create fake handler
func createfakeHandler() HandlerServer {
	request := httptest.NewRequest("GET", "http://link.com", nil)
	writer := httptest.NewRecorder()

	return HandlerServer{
		Writer:  writer,
		Request: request,
	}
}

// Create fake articles
func fakeArticles() []model.Article {
	articles := model.InitArticles(1, "First title", "First description")
	articles = model.AddArticleItem(articles, 2, "Second title", "Second description")
	return articles
}

// Create new context with CSV
func TestNewContextCSV(t *testing.T) {
	var handlerServer HandlerServer = createfakeHandler()
	var articles []model.Article = fakeArticles()
	var csv ExportArticles = CSV{}

	context := NewContext(csv, articles, handlerServer)

	t.Run("Articles not empty", func(t *testing.T) {
		if len(context.Articles) <= 0 {
			t.Errorf("articles = %v, want 2", len(context.Articles))
		}
	})

	var emptyArticles []model.Article = model.EmptyArticles()
	contextEmptyArticles := NewContext(csv, emptyArticles, handlerServer)

	t.Run("Articles empty", func(t *testing.T) {
		if len(contextEmptyArticles.Articles) > 0 {
			t.Errorf("articles = %v, want 0", len(contextEmptyArticles.Articles))
		}
	})
}

// Create new context with HLSX
func TestNewContextHLSX(t *testing.T) {
	var handlerServer HandlerServer = createfakeHandler()
	var articles []model.Article = fakeArticles()
	var xlsx ExportArticles = XLSX{}

	context := NewContext(xlsx, articles, handlerServer)

	t.Run("Articles not empty", func(t *testing.T) {
		if len(context.Articles) <= 0 {
			t.Errorf("articles = %v, want 2", len(context.Articles))
		}
	})

	var emptyArticles []model.Article = model.EmptyArticles()
	contextEmptyArticles := NewContext(xlsx, emptyArticles, handlerServer)

	t.Run("Articles empty", func(t *testing.T) {
		if len(contextEmptyArticles.Articles) > 0 {
			t.Errorf("articles = %v, want 0", len(contextEmptyArticles.Articles))
		}
	})
}

// Make csv export
func TestMakeExportCSV(t *testing.T) {
	var handlerServer HandlerServer = createfakeHandler()
	var articles []model.Article = fakeArticles()
	var csv ExportArticles = CSV{}

	context := NewContext(csv, articles, handlerServer)
	err := context.MakeExport()

	t.Run("Get error", func(t *testing.T) {
		if err != nil {
			t.Errorf("error = %v, want nil", err)
		}
	})
}

// Make xlsx export
func TestMakeExportHLSX(t *testing.T) {
	var handlerServer HandlerServer = createfakeHandler()
	var articles []model.Article = fakeArticles()
	var xlsx ExportArticles = XLSX{}

	context := NewContext(xlsx, articles, handlerServer)
	err := context.MakeExport()

	t.Run("Get error", func(t *testing.T) {
		if err != nil {
			t.Errorf("error = %v, want nil", err)
		}
	})
}
