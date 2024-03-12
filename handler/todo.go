package handler

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo/internal/database"
	"todo/model"
	utils "todo/util"
)

func (cfg *Config) CreateTODO(writer http.ResponseWriter, request *http.Request, currentUser database.User) {
	todoName := request.FormValue("todoName")
	todoParams := database.CreateTodoParams{
		ID:        uuid.New(),
		Name:      sql.NullString{String: todoName, Valid: true},
		Owner:     currentUser.ID,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	cfg.DB.CreateTodo(request.Context(), todoParams)
	http.Redirect(writer, request, "/home", 301)
}

func (cfg *Config) GetTODO(writer http.ResponseWriter, request *http.Request, currentUser database.User) {
	todoId := chi.URLParam(request, "todoId")
	todo, _ := cfg.DB.FindTodoById(request.Context(), uuid.MustParse(todoId))

	tasks, _ := cfg.DB.GetTasks(request.Context(), todo.ID)

	modelTodo := model.Todo{
		Id:        todo.ID,
		Name:      todo.Name.String,
		CreatedAt: todo.CreatedAt.Time,
		UpdatedAt: todo.UpdatedAt.Time,
		Tasks:     utils.MapManyTasksFromDB(tasks),
	}
	utils.RenderTemplate(writer, modelTodo, "template/todo-details.html")
}
