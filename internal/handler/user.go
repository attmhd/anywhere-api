package handlers

import (
	"anywhere-api/internal/models"
	"anywhere-api/internal/services"
	"anywhere-api/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UserHandler mendefinisikan struct yang berisi service untuk operasi user
type UserHandler struct {
	service services.UserService
}

// NewUserHandler menginisialisasi handler dengan service
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Authenticate the user and generate JWT token
	token, err := h.service.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return helper.ErrorResponse(c, "Authentication failed", err.Error())
	}

	// Return the JWT token in the response
	return helper.SuccessResponse(c, "Login successful", fiber.Map{
		"token": token,
	})
}

// CreateUser menangani HTTP request untuk membuat user baru
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.ErrorResponse(c, "Failed to hash password", err.Error())
	}

	user.Password = hashedPassword

	if err := h.service.CreateUser(&user); err != nil {
		return helper.ErrorResponse(c, "Failed to create user", err.Error())
	}

	return helper.SuccessResponse(c, "User created successfully", user)
}

// GetUserByID menangani HTTP request untuk mengambil user berdasarkan ID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return helper.ErrorResponse(c, "Failed to get user", err.Error())
	}

	if user == nil {
		return helper.ErrorResponse(c, "User not found", "User not found")
	}

	return helper.SuccessResponse(c, "User found", user)
}

// GetUserByUsername menangani HTTP request untuk mencari user berdasarkan username
func (h *UserHandler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		return helper.ErrorResponse(c, "Failed to get user", err.Error())
	}

	if user == nil {
		return helper.ErrorResponse(c, "User not found", "User not found")
	}

	return helper.SuccessResponse(c, "User found", user)
}

// GetAllUsers menangani HTTP request untuk mengambil semua user
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return helper.ErrorResponse(c, "Failed to get users", err.Error())
	}

	return helper.SuccessResponse(c, "Users found", users)
}

// UpdateUser menangani HTTP request untuk memperbarui user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.service.UpdateUser(&user); err != nil {
		return helper.ErrorResponse(c, "Failed to update user", err.Error())
	}

	return helper.SuccessResponse(c, "User updated successfully", user)
}

// DeleteUser menangani HTTP request untuk menghapus user berdasarkan ID
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		return helper.ErrorResponse(c, "Failed to delete user", err.Error())
	}

	return helper.SuccessResponse(c, "User deleted successfully", nil)
}
