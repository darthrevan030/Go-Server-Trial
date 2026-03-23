package config

import (
	"log"
	"os"
)

type Config struct {
	MongoURI            string
	MongoDBName         string
	MongoCollectionName string
	Port                string
}

func Load() *Config {
	cfg := &Config{
		MongoURI:            getEnv("MONGODB_URI", ""),
		MongoDBName:         getEnv("MONGODB_NAME", "go-server-trial"),
		MongoCollectionName: getEnv("MONGODB_COLLECTION_NAME", "users"),
		Port:                getEnv("PORT", "3000"),
	}

	if cfg.MongoURI == "" {
		log.Fatal("MONGODB_URI is required")
	}

	return cfg
}

// getEnv reads an env var, falls back to a default if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
