package main

import (
	"context"
	"github.com/SemmiDev/Go-Scheduled/config"
	"github.com/SemmiDev/Go-Scheduled/controller"
	"github.com/SemmiDev/Go-Scheduled/repository"
	"github.com/SemmiDev/Go-Scheduled/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	log.Println("Reading the configuration from environment variables ...")
	config.App()

	log.Println("Initializing the database connection ...")
	conn, err := pgxpool.Connect(context.Background(), config.PostgreURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("Initializing the repository ...")
	assignmentRepo := repository.New(conn)

	log.Println("Initializing the service ...")
	assignmentService := service.NewTaskService(assignmentRepo)

	log.Println("Initializing the controller ...")
	assignmentController := controller.NewTaskController(assignmentService)

	log.Println("Initializing the server ...")
	app := fiber.New(config.Fiber())

	log.Println("Initializing the routes ...")
	assignmentController.SetupRoute(app)

	log.Println("Everything is OK ...")
	if err := app.Listen(config.AppPort); err != nil {
		panic(err)
	}
}
