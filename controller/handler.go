package controller

import (
	"database/sql"
	"errors"
	"github.com/SemmiDev/Go-Scheduled/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (cr *TaskController) CreateTaskHandler(c *fiber.Ctx) error {
	var request model.CreateTaskReq
	err := c.BodyParser(&request)
	if err != nil {
		return model.ApiResponse(c, http.StatusBadRequest, nil, err)
	}

	if err := request.Validate(); err != nil {
		return model.ApiResponse(c, http.StatusBadRequest, nil, err)
	}

	var response *model.Result
	response, err = cr.service.CreateTask(c.Context(), &request)
	if err != nil {
		if errors.Is(err, model.ErrParseTime) {
			return model.ApiResponse(c, http.StatusBadRequest, nil, err)
		}
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}

	return model.ApiResponse(c, http.StatusCreated, response, nil)
}

var ErrEmptyID = errors.New("ID is empty")

func (cr *TaskController) UpdateTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ApiResponse(c, http.StatusBadRequest, nil, ErrEmptyID)
	}

	var request model.UpdateTaskReq
	err := c.BodyParser(&request)
	if err != nil {
		return model.ApiResponse(c, http.StatusBadRequest, nil, err)
	}

	if err := request.Validate(); err != nil {
		return model.ApiResponse(c, http.StatusBadRequest, nil, err)
	}

	request.ID = ID
	var response *model.Result
	response, err = cr.service.UpdateTask(c.Context(), &request)
	if err != nil {
		if strings.Contains(err.Error(), "parse time") {
			return model.ApiResponse(c, http.StatusBadRequest, nil, err)
		}
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}

	return model.ApiResponse(c, http.StatusOK, response, nil)
}

func (cr *TaskController) GetTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ApiResponse(c, http.StatusBadRequest, nil, ErrEmptyID)
	}

	response, err := cr.service.GetTaskById(c.Context(), ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ApiResponse(c, http.StatusNotFound, nil, err)
		}
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}

	return model.ApiResponse(c, http.StatusOK, response, nil)
}

func (cr *TaskController) GetListTasksHandler(c *fiber.Ctx) error {
	response, err := cr.service.GetListTasks(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ApiResponse(c, http.StatusNotFound, nil, err)
		}
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}

	return model.ApiResponse(c, http.StatusOK, response, nil)
}

func (cr *TaskController) MarkAsCompletedTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ApiResponse(c, http.StatusBadRequest, nil, ErrEmptyID)
	}

	response, err := cr.service.UpdateStatusTaskById(c.Context(), ID)
	if err != nil {
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}
	return model.ApiResponse(c, http.StatusOK, response, nil)
}

func (cr *TaskController) DeleteTaskHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	if ID == "" {
		return model.ApiResponse(c, http.StatusBadRequest, nil, ErrEmptyID)
	}

	if err := cr.service.DeleteTask(c.Context(), ID); err != nil {
		return model.ApiResponse(c, http.StatusInternalServerError, nil, err)
	}
	return c.SendStatus(http.StatusOK)
}
