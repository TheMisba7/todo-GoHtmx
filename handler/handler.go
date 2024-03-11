package handler

import (
	"github.com/google/uuid"
	"net/http"
	"todo/internal/database"
)

type Config struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (config *Config) HomePage(writer http.ResponseWriter, request *http.Request, user database.User) {
	writer.Write([]byte("welcome to home page"))
}

func (config *Config) Middleware(next authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("userId")
		if err != nil {
			// cookie not found
			http.Redirect(writer, request, "/login", 301)
			return
		}

		userById, err := config.DB.GetUserById(request.Context(), uuid.MustParse(cookie.Value))
		if err != nil {
			panic(err)
		}
		next(writer, request, userById)
	}
}
