package services

import (
	"net/http"
)

// Make HTTP call to go in route page
func Redirect(w http.ResponseWriter, r *http.Request, route string) {
	http.Redirect(w, r, route, http.StatusSeeOther)
}
