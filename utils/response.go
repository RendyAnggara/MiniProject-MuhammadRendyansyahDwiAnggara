package utils

import "github.com/gofiber/fiber/v3"

type Response struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
	Data    any      `json:"data"`
}

func SuccessResponse(c fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

func ErrorResponse(c fiber.Ctx, statusCode int, message string, errors []string) error {
	return c.Status(statusCode).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}