package user

import (
	"strconv"

	"bangkho.dev/tcl/test/backend/helpers"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles user HTTP requests
type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// userHandler implements UserHandler
type userHandler struct {
	service UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(service UserService) UserHandler {
	return &userHandler{service: service}
}

// Register handles POST /users/register
func (h *userHandler) Register(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.BadRequest(c, "Invalid request body")
	}

	if errs := helpers.ValidateStruct(req); len(errs) > 0 {
		return helpers.BadRequestWithErrors(c, "Validation failed", errs)
	}

	result, err := h.service.Register(&req)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Created(c, "User registered successfully", result.ToResponse())
}

// Login handles POST /users/login
func (h *userHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.BadRequest(c, "Invalid request body")
	}

	if errs := helpers.ValidateStruct(req); len(errs) > 0 {
		return helpers.BadRequestWithErrors(c, "Validation failed", errs)
	}

	result, err := h.service.Login(&req)
	if err != nil {
		return helpers.Unauthorized(c, "Invalid username or password")
	}

	return helpers.Success(c, "Login successful", result)
}

// GetByID handles GET /users/:id
func (h *userHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid user ID")
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		return helpers.NotFound(c, "User not found")
	}

	return helpers.Success(c, "User retrieved successfully", result.ToResponse())
}

// GetAll handles GET /users
func (h *userHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	result, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Users retrieved successfully", result)
}

// Update handles PUT /users/:id
func (h *userHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid user ID")
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.BadRequest(c, "Invalid request body")
	}

	if errs := helpers.ValidateStruct(req); len(errs) > 0 {
		return helpers.BadRequestWithErrors(c, "Validation failed", errs)
	}

	result, err := h.service.Update(uint(id), &req)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "User updated successfully", result.ToResponse())
}

// Delete handles DELETE /users/:id
func (h *userHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid user ID")
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "User deleted successfully", nil)
}
