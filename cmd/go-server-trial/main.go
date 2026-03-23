package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/darthrevan030/go-server-trial/internal/mongodb"
	"github.com/darthrevan030/go-server-trial/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error While Reading Environment Variables")
	}

	slog.Info("Environment Variables Loaded Successfully")
}


func mongoConnectionOpen() *mongo.Client {
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(opts)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB: ", err)
    }

    // ping to verify
    var result bson.M
    if err := client.Database(os.Getenv("MONGODB_URI")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
        log.Fatal("MongoDB ping failed: ", err)
    }

    log.Println("Connected to MongoDB!")
    return client
}


func main() {

	client := mongoConnectionOpen()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	} ()

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
		r.Put("/users/{id}", UserService.UpdateUserAgeByID)
		r.Delete("/users/{id}", UserService.DeleteUserByID)
		r.Delete("/users", UserService.DeleteAllUsers)
	
	})
	log.Println("Server running on :3000")  
    log.Fatal(http.ListenAndServe(":3000", r))
	http.ListenAndServe(":3000", r)
}