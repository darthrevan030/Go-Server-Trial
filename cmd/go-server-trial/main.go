package main

import (
	"net/http"

	"github.com/darthrevan030/go-server-trial/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// userservice instance
	UserService := service.UserService{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Server is healthy"))
		})
	
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Welcome"))	
		})

		r.Post("/users", UserService.CreateUser)
		r.Get("/users", UserService.GetAllUsers)
		r.Get("/users/{id}", UserService.GetUserByID)
		r.Put("/users{id}", UserService.UpdateUserAgeByID)
		r.Delete("/users/{id}", UserService.DeleteUserByID)
		r.Delete("/users", UserService.DeleteAllUsers)
	
	})

	http.ListenAndServe(":3000", r)
}