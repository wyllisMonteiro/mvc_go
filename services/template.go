package services

import (
	"text/template"
	"net/http"
)

type Render struct {
	ParseFiles []string
	Writter http.ResponseWriter
	Data interface{}
}

func RenderTemplate(render Render) (error){
	tmpl, err := template.ParseFiles(render.ParseFiles...)
	if err != nil {
		return err
	}

	err = tmpl.Execute(render.Writter, render.Data)
	if err != nil {
		return err
	}

	return nil
}