package controller

import (
	"github.com/SemmiDev/Go-Scheduled/service"
	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (cr *TaskController) SetupRoute(app *fiber.App) {
	app.Post("/api/tasks", cr.CreateTaskHandler)
	app.Get("/api/tasks", cr.GetListTasksHandler)
	app.Get("/api/tasks/:id", cr.GetTaskHandler)
	app.Put("/api/tasks/:id", cr.UpdateTaskHandler)
	app.Put("/api/tasks/:id/status", cr.MarkAsCompletedTaskHandler)
	app.Delete("/api/tasks/:id", cr.DeleteTaskHandler)
}
