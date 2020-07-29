package export

import (
    "os"
    "encoding/csv"
    "time"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type CSV struct {}

func (exportCSV CSV) ExportAsFile(datas []model.Article) {
	file, err := os.OpenFile("assets/csv/article_" + formatCurrentDate() + ".csv", os.O_CREATE|os.O_WRONLY, 0777)
     
    if err != nil {
        return
    }

    defer file.Close()

    init_array := []string{"Titre", "Description"}

    strWrite := [][]string{init_array}
    for _, item := range datas {
        strWrite = append(strWrite, []string{item.Title, item.Description})
    }

    csvWriter := csv.NewWriter(file)

    defer csvWriter.Flush()

    err = csvWriter.WriteAll(strWrite)
    if err != nil {
        return
    }
}

func formatCurrentDate() (string) {
    today := time.Now()
    return today.Format("2006.01.02 15:04")
} 