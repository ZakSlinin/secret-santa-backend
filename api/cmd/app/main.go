package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/secret-santa-rubot/api/internal/handler"
	"github.com/secret-santa-rubot/api/internal/repository"
	"github.com/secret-santa-rubot/api/internal/service"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// DB connection
	db, err := ConnectionDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// User endpoint setup
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Post("/api/register", userHandler.RegisterUser)
	http.ListenAndServe(":8080", r)
}

func ConnectionDB() (*sql.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/santa?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
