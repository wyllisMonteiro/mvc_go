package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type Context struct {
	exportArticles ExportArticles
	articles []model.Article
}

func NewContext(e ExportArticles, a []model.Article) *Context {
	return &Context{
		exportArticles: e,
		articles: a,
	}
}

func (e *Context) MakeExport() {
    e.exportArticles.ExportAsFile(e.articles)
}