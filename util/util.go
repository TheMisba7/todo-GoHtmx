package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"todo/internal/database"
	"todo/model"
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

func MapManyTasksFromDB(tasks []database.Task) []model.Task {
	modelTasks := make([]model.Task, len(tasks))
	for i, task := range tasks {
		modelTasks[i] = MapOneTaskFromDB(task)
	}
	return modelTasks
}
func MapOneTaskFromDB(task database.Task) model.Task {
	return model.Task{
		Name:      task.Name.String,
		CreatedAt: task.CreatedAt.Time,
		UpdatedAt: task.UpdatedAt.Time,
		TodoId:    task.TodoID,
		Status:    int8(task.Status),
		StartDate: task.StartDate.Time,
		EndDate:   task.EndDate.Time,
	}
}

func MapOneTodoFromDB(todo database.Todo) model.Todo {
	return model.Todo{
		Name:      todo.Name.String,
		CreatedAt: todo.CreatedAt.Time,
		UpdatedAt: todo.UpdatedAt.Time,
		Id:        todo.ID,
		Owner:     todo.Owner,
	}
}
