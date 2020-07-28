package export

import "fmt"

func GetExport(exportType string) (iExport, error) {
    if exportType == "csv" {
        return NewCSV(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}