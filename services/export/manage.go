package export

import (
  model "github.com/wyllisMonteiro/go_mvc/models"
)

func ManageExport(type_export string, articles []model.Article) {
	var export_file *Export

	switch type_export {
	case "csv":
		csv := &CSV{}
		export_file = InitExport(csv, articles)
	case "xlsx":
		xlsx := &XLSX{}
		export_file = InitExport(xlsx, articles)
	default:
		return
	}

	export_file.MakeExport()
}