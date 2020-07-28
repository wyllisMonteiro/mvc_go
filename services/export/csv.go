package export

import (
    "os"
    "fmt"
    "encoding/csv"
    "time"
	model "github.com/wyllisMonteiro/go_mvc/models"
)

type t_csv struct {
	export
}

func NewCSV(datas []model.Article) (error) {
    date_format := formatCurrentDate()
	file, err := os.OpenFile("article_" + date_format + ".csv", os.O_CREATE|os.O_WRONLY, 0777)
	
    defer file.Close()
 
    if err != nil {
        return err
    }

    init_array := []string{"Titre", "Description"}

    strWrite := [][]string{init_array}
    for _, item := range datas {
        strWrite = append(strWrite, []string{item.Title, item.Description})
    }

    csvWriter := csv.NewWriter(file)
    csvWriter.WriteAll(strWrite)
	csvWriter.Flush()

	return nil
}

func formatCurrentDate() (string) {
    today := time.Now()
    return today.Format("2006.01.02 15:04")
} 