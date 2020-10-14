package export

import (
	"encoding/csv"
	"net/http"
	"time"

	model "github.com/wyllisMonteiro/warticle/models"
)

// Data about CSV
type CSV struct{}

// Export articles as CSV file from browser
// Go in context.go file to understand HandlerServer
func (exportCSV CSV) ExportAsFile(datas []model.Article, handler HandlerServer) error {
	contentDisposition := "attachment; filename=articles_" + formatCurrentDate() + ".csv"
	handler.Writer.Header().Set("Content-Disposition", contentDisposition)
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

// Get current date as string with format "YYYY.MM.DD H:i"
func formatCurrentDate() string {
	today := time.Now()
	return today.Format("2006.01.02 15:04")
}
