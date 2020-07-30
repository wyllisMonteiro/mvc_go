package services

import (
	"net/http"
	"text/template"
)

/**
 *
 * Render struct is all data needed to display a page
 *
 * ParseFiles is all templates (/web folder) needed to display page
 * 		> Should contain web/base.tmpl file
 *		> Put at index 0 of array page that overrides web/base.tmpl like web/articles.tmpls
 *		> Example : parseFiles := []string{"web/articles.tmpl", "web/base.tmpl"}
 * Writer is the response of http request
 * Data is dynamic data to pass to template like articles ...
 *
 */
type Render struct {
	ParseFiles []string
	Writer     http.ResponseWriter
	Data       interface{}
}

/**
 *
 * Display a page in a browser
 *
 */
func RenderTemplate(render Render) error {
	tmpl, err := template.ParseFiles(render.ParseFiles...)
	if err != nil {
		return err
	}

	err = tmpl.Execute(render.Writer, render.Data)
	if err != nil {
		return err
	}

	return nil
}
