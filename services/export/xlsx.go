package export

import (
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type XLSX struct {}

func (exportXLSX XLSX) ExportAsFile(datas []model.Article) {
	f := excelize.NewFile()
    // Set value of a cell.
    f.SetCellValue("Sheet1", "A1", "Titre")
	f.SetCellValue("Sheet1", "B1", "Description")
	
	for index, item := range datas {
		// first index = 0 and does not exit in excel
		// second index = 1 is already use before 
		index_excel := strconv.Itoa(index + 2)
        f.SetCellValue("Sheet1", "A" + index_excel, item.Title)
        f.SetCellValue("Sheet1", "B" + index_excel, item.Description)
    }

    // Save xlsx file by the given path.
    if err := f.SaveAs("assets/xlsx/articles_" + formatCurrentDate() + ".xlsx"); err != nil {
		println(err.Error())
		return
	}

	return
}