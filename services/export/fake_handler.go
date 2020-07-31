package export

import (
	"net/http"
	"net/http/httptest"
)

// CreateFakeHandler() : create fake handler to get http.ResponseWriter, *http.Request
func CreateFakeHandler() (http.ResponseWriter, *http.Request) {
	request := httptest.NewRequest("GET", "http://link.com", nil)
	writer := httptest.NewRecorder()

	return writer, request
}
