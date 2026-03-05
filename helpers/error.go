package helpers

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Response represents a standard success response structure
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// BadRequest returns a 400 Bad Request error
func BadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: message,
	})
}

// BadRequestWithErrors returns a 400 Bad Request error with validation errors
func BadRequestWithErrors(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: message,
		Errors:  errors,
	})
}

// Unauthorized returns a 401 Unauthorized error
func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	})
}

// Forbidden returns a 403 Forbidden error
func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(ErrorResponse{
		Code:    fiber.StatusForbidden,
		Message: message,
	})
}

// NotFound returns a 404 Not Found error
func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
		Code:    fiber.StatusNotFound,
		Message: message,
	})
}

// Conflict returns a 409 Conflict error
func Conflict(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(ErrorResponse{
		Code:    fiber.StatusConflict,
		Message: message,
	})
}

// InternalServerError returns a 500 Internal Server Error
func InternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	})
}

// Success returns a standard success response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Code:    fiber.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Created returns a 201 Created response
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Code:    fiber.StatusCreated,
		Message: message,
		Data:    data,
	})
}
