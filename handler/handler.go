package handler

import (
	"github.com/google/uuid"
	"net/http"
	"todo/internal/database"
	utils "todo/util"
)

type Config struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) HomePage(writer http.ResponseWriter, request *http.Request, user database.User) {
	todos, _ := cfg.DB.FindByUser(request.Context(), user.ID)
	utils.RenderTemplate(writer, todos, "template/home.html")
}

func (cfg *Config) PostLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username := request.FormValue("username")
		password := request.FormValue("password")
		user, err := cfg.DB.GetUser(request.Context(), username)
		if err != nil {
			panic(err)
		}
		if utils.CheckPasswordHash(password, user.Password) {
			http.SetCookie(writer, &http.Cookie{Name: "userId", Value: user.ID.String()})
			http.Redirect(writer, request, "/home", 301)
		}
		http.Redirect(writer, request, "/login", 301)
	}
}

func (cfg *Config) CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username := request.FormValue("username")
		password := request.FormValue("password")
		userParams := database.CreateUserParams{
			ID:       uuid.New(),
			Username: username,
			Password: utils.HashPassword(password),
		}
		user, err := cfg.DB.CreateUser(request.Context(), userParams)
		if err != nil {
			panic(err)
		}
		http.SetCookie(writer, &http.Cookie{Name: "userId", Value: user.ID.String()})
		http.Redirect(writer, request, "/home", 301)
	}
}

func (cfg *Config) Middleware(next authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("userId")
		if err != nil {
			// cookie not found
			http.Redirect(writer, request, "/login", 301)
			return
		}

		userById, err := cfg.DB.GetUserById(request.Context(), uuid.MustParse(cookie.Value))
		if err != nil {
			panic(err)
		}
		next(writer, request, userById)
	}
}
