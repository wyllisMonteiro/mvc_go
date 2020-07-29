package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type Export struct {
	exportArticles ExportArticles
	articles []model.Article
}

func InitExport(e ExportArticles, a []model.Article) *Export {
	export_file := &Export{
		exportArticles: e,
		articles: a,
	}

	MakeExport(export_file)
	return export_file
}

func MakeExport(e *Export) {
    e.exportArticles.ExportAsFile(e.articles)
}