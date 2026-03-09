package main

import (
	"context"
	"time"
	"log"
	"net/http"

	"github.com/darthrevan030/go-server-trial/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnectionOpen() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/go-server-trial"))

	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}

	log.Println("Connected to MongoDB")
	return client
} 


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