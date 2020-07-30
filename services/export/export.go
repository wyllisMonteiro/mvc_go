package export

import (
	"net/http"

	model "github.com/wyllisMonteiro/go_mvc/models"
)

type HandlerServer struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type Context struct {
	exportArticles ExportArticles
	articles       []model.Article
	handlerServer  HandlerServer
}

func NewContext(e ExportArticles, a []model.Article, h HandlerServer) *Context {
	return &Context{
		exportArticles: e,
		articles:       a,
		handlerServer:  h,
	}
}

func (e *Context) MakeExport() error {
	err := e.exportArticles.ExportAsFile(e.articles, e.handlerServer)
	if err != nil {
		return err
	}

	return nil
}
