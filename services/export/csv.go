package export

import (
	"encoding/csv"
	"net/http"
	"time"

	model "github.com/wyllisMonteiro/go_mvc/models"
)

type CSV struct{}

func (exportCSV CSV) ExportAsFile(datas []model.Article, handler HandlerServer) error {
	handler.Writer.Header().Set("Content-Disposition", "attachment; filename=export.csv")
	handler.Writer.Header().Set("Content-Type", "text/csv")
	handler.Writer.Header().Set("Transfer-Encoding", "chunked")

	writer := csv.NewWriter(handler.Writer)
	err := writer.Write([]string{"Titre", "Description"})
	if err != nil {
		http.Error(handler.Writer, "cannot write to file", http.StatusInternalServerError)
		return err
	}

	for _, item := range datas {
		err := writer.Write([]string{item.Title, item.Description})
		if err != nil {
			http.Error(handler.Writer, "cannot write to file", http.StatusInternalServerError)
			return err
		}
	}

	writer.Flush()

	return nil
}

func formatCurrentDate() string {
	today := time.Now()
	return today.Format("2006.01.02 15:04")
}
