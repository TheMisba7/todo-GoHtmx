package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"todo/handler"
	"todo/internal/database"
	utils "todo/util"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	dbAddr := os.Getenv("DATASOURCE")
	db, err := sql.Open("postgres", dbAddr)
	if db == nil {
		panic("db connection cannot be nil")
	}
	config := handler.Config{DB: database.New(db)}
	defer db.Close()
	mainRouter := chi.NewRouter()

	mainRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		utils.RenderTemplate(writer, nil, "template/index.html")
	})

	mainRouter.Get("/home", config.Middleware(config.HomePage))

	mainRouter.Get("/login", func(writer http.ResponseWriter, request *http.Request) {
		utils.RenderTemplate(writer, nil, "template/login.html")
	})
	mainRouter.Post("/login", config.PostLogin())

	mainRouter.Post("/register", config.CreateUser())
	mainRouter.Get("/register", func(writer http.ResponseWriter, request *http.Request) {
		utils.RenderTemplate(writer, nil, "template/register.html")
	})

	mainRouter.Get("/todo/new", func(writer http.ResponseWriter, request *http.Request) {
		utils.RenderTemplate(writer, nil, "template/create-todo.html")
	})
	mainRouter.Post("/todo", config.Middleware(config.CreateTODO))
	mainRouter.Get("/todo/{todoId}", config.Middleware(config.GetTODO))

	mainRouter.Post("/task", config.Middleware(config.CreateTask))
	server := http.Server{
		Handler: mainRouter,
		Addr:    fmt.Sprintf(":%v", port),
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
