package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
