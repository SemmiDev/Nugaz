package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port       string
	DBDriver   string
	PostgreURL string
)

type Config interface {
	Get(key string) string
}

type config struct{}

func (c *config) Get(key string, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	log.Println(key + " :: default ")
	return def
}

func App() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	c := new(config)
	Port = c.Get("PORT", ":9090")
	PostgreURL = c.Get("POSTGRE_URL", "postgres://sammidev:sammidev@localhost:5432/tasks?sslmode=disable")
	DBDriver = c.Get("DB_DRIVER", "postgres")
}
