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

	
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Server is healthy"))
	})
	
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome"))	
	})

	router.Post("api/v1/users", )
	
	http.ListenAndServe(":3000", router)
}