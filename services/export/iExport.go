package export

import (
	model "github.com/wyllisMonteiro/go_mvc/models"
)

// ExportArticles : Export all articles from browser
type ExportArticles interface {
	ExportAsFile(datas []model.Article, handler HandlerServer) error
}
