package model

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func ApiResponse(c *fiber.Ctx, status int, data interface{}, err error) error {
	c.Status(status)
	return c.JSON(response{
		Code:   status,
		Status: http.StatusText(status),
		Data:   data,
		Error:  err.Error(),
	})
}
