package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	AppPort    string
	PostgreURL string
)

type config struct{}

func (c *config) Get(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func App() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("config err: %v", err)
	}

	c := new(config)
	AppPort = c.Get("APP_PORT", ":9090")
	PostgreURL = c.Get("POSTGRE_URL", "postgres://sammidev:sammidev@localhost:5432/tasks?sslmode=disable")
}
