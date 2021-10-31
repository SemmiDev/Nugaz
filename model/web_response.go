package model

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type SuccessResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Error struct {
	Msg string `json:"error_msg"`
}

func ErrorResponse(c *fiber.Ctx, status int, err error) error {
	c.Status(status)
	return c.JSON(SuccessResponse{
		Code:   status,
		Status: http.StatusText(status),
		Data: Error{
			Msg: err.Error(),
		},
	})
}
