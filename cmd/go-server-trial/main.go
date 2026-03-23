package main

import (
	"context"
	"log"
	"net/http"

	"github.com/darthrevan030/go-server-trial/internal/config"
	"github.com/darthrevan030/go-server-trial/internal/database"
	"github.com/darthrevan030/go-server-trial/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := config.Load()

	mongoClient := database.MongodbConnect(cfg.MongoURI)
	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database(cfg.MongoDBName)

	userRepo := user.NewRepository(db, cfg.MongoCollectionName)
	userHandler := user.NewHandler(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Server is healthy"))
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome"))
		})

		userHandler.RegisterRoutes(r)
	})

	log.Println("Server running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
