package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, data any, path string) {
	tmp := template.Must(template.ParseFiles(path))
	err := tmp.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
