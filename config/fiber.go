package config

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Fiber() fiber.Config {
	return fiber.Config{
		AppName:      "Nugaz",
		BodyLimit:    2000,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
