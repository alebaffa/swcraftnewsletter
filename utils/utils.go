package utils

import (
	"html/template"
	"net/http"
)

func ExecuteTemplate(templateType string, res http.ResponseWriter) {
	templateType = templateType + ".html"
	templates := template.Must(template.ParseGlob("assets/*.html"))
	if err := templates.ExecuteTemplate(res, templateType, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
