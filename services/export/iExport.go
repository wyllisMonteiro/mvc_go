package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type ExportArticles interface {
	ExportAsFile(datas []model.Article, handler HandlerServer) error
}
