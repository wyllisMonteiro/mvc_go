package services

import (
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request, route string) {
	http.Redirect(w, r, route, http.StatusSeeOther)
}