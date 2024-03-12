package handler

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo/internal/database"
	utils "todo/util"
)

func (cfg *Config) CreateTask(writer http.ResponseWriter, request *http.Request, currentUser database.User) {
	taskName := request.FormValue("taskName")
	todoId := request.FormValue("todoId")
	task := database.CreateTaskParams{
		ID:        uuid.New(),
		TodoID:    uuid.MustParse(todoId),
		Name:      sql.NullString{String: taskName, Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	_, err2 := cfg.DB.CreateTask(request.Context(), task)
	if err2 != nil {
		panic(err2)
	}

	todo, _ := cfg.DB.FindTodoById(request.Context(), uuid.MustParse(todoId))
	tasks, err := cfg.DB.GetTasks(request.Context(), todo.ID)
	if err != nil {
		return
	}

	modelTodo := utils.MapOneTodoFromDB(todo)
	modelTodo.Tasks = utils.MapManyTasksFromDB(tasks)
	utils.RenderTemplate(writer, modelTodo, "template/todo-details.html")
}

func (cfg *Config) deleteTask(writer http.ResponseWriter, request *http.Request, currentUser database.User) {
	taskId := chi.URLParam(request, "taskId")
	task, err := cfg.DB.GetTaskById(request.Context(), uuid.MustParse(taskId))
	if err != nil {
		panic(err)
	}
	err = cfg.DB.DeleteTask(request.Context(), task.ID)
	if err != nil {
		panic(err)
	}

	todo, _ := cfg.DB.FindTodoById(request.Context(), task.TodoID)
	tasks, _ := cfg.DB.GetTasks(request.Context(), todo.ID)
	modelTodo := utils.MapOneTodoFromDB(todo)
	modelTodo.Tasks = utils.MapManyTasksFromDB(tasks)
	utils.RenderTemplate(writer, todo, "template/todo-details.html")
}
