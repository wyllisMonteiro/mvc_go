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
	FileType ExportArticles
	Articles []model.Article
	Handler  HandlerServer
}

// NewContext() : Init new context
func NewContext(e ExportArticles, a []model.Article, h HandlerServer) *Context {
	return &Context{
		FileType: e,
		Articles: a,
		Handler:  h,
	}
}

// MakeExport() : Make export according to the context (CSV or XLSX)
func (e *Context) MakeExport() error {
	err := e.FileType.ExportAsFile(e.Articles, e.Handler)
	if err != nil {
		return err
	}

	return nil
}
