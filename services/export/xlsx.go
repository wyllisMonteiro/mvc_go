package export

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type XLSX struct{}

func (exportXLSX XLSX) ExportAsFile(datas []model.Article, handler HandlerServer) error {
	contentDisposition := "attachment; filename=articles_" + formatCurrentDate() + ".xlsx"
	handler.Writer.Header().Set("Content-Disposition", contentDisposition)
	handler.Writer.Header().Set("Content-Type", "application/vnd.ms-excel")
	handler.Writer.Header().Set("Transfer-Encoding", "chunked")

	xlsx := excelize.NewFile()
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A1", "Titre")
	xlsx.SetCellValue("Sheet1", "B1", "Description")

	for index, item := range datas {
		// first index = 0 and does not exit in excel
		// second index = 1 is already use before
		index_excel := strconv.Itoa(index + 2)
		xlsx.SetCellValue("Sheet1", "A"+index_excel, item.Title)
		xlsx.SetCellValue("Sheet1", "B"+index_excel, item.Description)
	}

	err := xlsx.Write(handler.Writer)
	if err != nil {
		return err
	}

	return nil
}
