package models

import (
	"strings"
	"testing"
)

func TestConnectToBDD(t *testing.T) {
	tests := []struct {
		name         string
		desscription string
		want         string
	}{
		{
			desscription: "Is there errors",
			want:         "No errors in err",
		},
		{
			name:         "error",
			desscription: "Is there errors",
			want:         "Database configuration error",
		},
	}

	for index, oneTest := range tests {
		t.Run(oneTest.name, func(t *testing.T) {
			if index == 1 {
				USER_DB = "test" // FALSE
			}

			_, err := ConnectToBDD()
			if err != nil {
				// Database configuration error
				isBadDatabaseConfig := strings.Contains(err.Error(), "Access denied for user")
				if isBadDatabaseConfig {
					return
				}

				t.Errorf("ConnectToBDD() = %v, want %v", err.Error(), oneTest.want)
			}
		})
	}
}
