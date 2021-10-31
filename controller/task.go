package controller

import (
	"database/sql"
	"errors"
	"github.com/SemmiDev/Go-Scheduled/model"
	"github.com/SemmiDev/Go-Scheduled/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
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

func (cr *TaskController) CreateTaskHandler(c *fiber.Ctx) error {
	var request model.CreateTaskReq
	err := c.BodyParser(&request)
	if err != nil {
		return model.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := request.Validate(); err != nil {
		return model.ErrorResponse(c, http.StatusBadRequest, err)
	}

	var response *model.Result
	response, err = cr.service.CreateTask(c.Context(), &request)
	if err != nil {
		if strings.Contains(err.Error(), "parse time") {
			return model.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(model.SuccessResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   response,
	})
}

func (cr *TaskController) UpdateTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ErrorResponse(c, http.StatusBadRequest, errors.New("ID is empty"))
	}

	var request model.UpdateTaskReq
	err := c.BodyParser(&request)
	if err != nil {
		return model.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := request.Validate(); err != nil {
		return model.ErrorResponse(c, http.StatusBadRequest, err)
	}

	request.ID = ID
	var response *model.Result
	response, err = cr.service.UpdateTask(c.Context(), &request)
	if err != nil {
		if strings.Contains(err.Error(), "parse time") {
			return model.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(model.SuccessResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func (cr *TaskController) GetTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ErrorResponse(c, http.StatusBadRequest, errors.New("ID is empty"))
	}

	response, err := cr.service.GetTaskById(c.Context(), ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ErrorResponse(c, http.StatusNotFound, err)
		}
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(model.SuccessResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func (cr *TaskController) GetListTasksHandler(c *fiber.Ctx) error {
	response, err := cr.service.GetListTasks(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ErrorResponse(c, http.StatusNotFound, err)
		}
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(model.SuccessResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func (cr *TaskController) MarkAsCompletedTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ErrorResponse(c, http.StatusBadRequest, errors.New("ID is empty"))
	}

	response, err := cr.service.UpdateStatusTaskById(c.Context(), ID)
	if err != nil {
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(model.SuccessResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func (cr *TaskController) DeleteTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ErrorResponse(c, http.StatusBadRequest, errors.New("ID is empty"))
	}

	if err := cr.service.DeleteTask(c.Context(), ID); err != nil {
		return model.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.SendStatus(http.StatusOK)
}
