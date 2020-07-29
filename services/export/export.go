package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type Export struct {
	exportArticles ExportArticles
	articles []model.Article
}

func InitExport(e ExportArticles, a []model.Article) *Export {
	return &Export{
		exportArticles: e,
		articles: a,
	}
}

func (e *Export) MakeExport() {
    e.exportArticles.ExportAsFile(e.articles)
}