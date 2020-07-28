package export

import (
	"fmt"
)

type csv struct {
	export
}

func NewCSV() iExport {
    return &csv{}
}