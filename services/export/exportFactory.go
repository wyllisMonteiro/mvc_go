package export

import (
    "fmt"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

func GetExport(exportType string, datas []model.Article) (error) {
    if exportType == "csv" {
        return NewCSV(datas)
    }
    return fmt.Errorf("Wrong gun type passed")
}