package helper

import (
	"github.com/gofiber/fiber/v2"
)

// CustomResponse is a struct for structuring API responses
type CustomResponse struct {
	Status  string      `json:"status"`          // Status: "success" or "fail"
	Message string      `json:"message"`         // A message describing the response
	Data    interface{} `json:"data,omitempty"`  // The data returned in case of success (optional)
	Error   string      `json:"error,omitempty"` // Error message in case of failure (optional)
}

// SuccessResponse creates a successful response with a given message and optional data
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	response := CustomResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// ErrorResponse creates a failure response with a given message and optional error
func ErrorResponse(c *fiber.Ctx, message string, err string) error {
	response := CustomResponse{
		Status:  "fail",
		Message: message,
		Error:   err,
	}
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}
