package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type ExportArticles interface {
	Export(datas []model.Article) (error)
}

var TypeExport = map[string]ExportArticles{
	"csv": ExportCSV{},
	"xlsx": ExportXLSX{},
}