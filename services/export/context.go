package export

import (
	"net/http"

	model "github.com/wyllisMonteiro/go_mvc/models"
)

// HandlerServer : Current handler params useful to make export from browser
type HandlerServer struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// Context : Get context export
// Design pattern strategy
// ExportArticles is an interface in iExport implemented in csv.go and xlsx.go files
type Context struct {
	exportArticles ExportArticles
	articles       []model.Article
	handlerServer  HandlerServer
}

// NewContext() : Init new conext
func NewContext(e ExportArticles, a []model.Article, h HandlerServer) *Context {
	return &Context{
		exportArticles: e,
		articles:       a,
		handlerServer:  h,
	}
}

// MakeExport() : Make export according to the context (CSV or XLSX)
func (e *Context) MakeExport() error {
	err := e.exportArticles.ExportAsFile(e.articles, e.handlerServer)
	if err != nil {
		return err
	}

	return nil
}
