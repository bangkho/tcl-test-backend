package customer

import (
	"strconv"

	"bangkho.dev/tcl/test/backend/helpers"

	"github.com/gofiber/fiber/v2"
)

// CustomerHandler handles customer HTTP requests
type CustomerHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// customerHandler implements CustomerHandler
type customerHandler struct {
	service CustomerService
}

// NewCustomerHandler creates a new customer handler
func NewCustomerHandler(service CustomerService) CustomerHandler {
	return &customerHandler{service: service}
}

// Create handles POST /customers
func (h *customerHandler) Create(c *fiber.Ctx) error {
	var req CreateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.BadRequest(c, "Invalid request body")
	}

	if errs := helpers.ValidateStruct(req); len(errs) > 0 {
		return helpers.BadRequestWithErrors(c, "Validation failed", errs)
	}

	result, err := h.service.Create(&req)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Created(c, "Customer created successfully", result.ToResponse())
}

// GetByID handles GET /customers/:id
func (h *customerHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid customer ID")
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		return helpers.NotFound(c, "Customer not found")
	}

	return helpers.Success(c, "Customer retrieved successfully", result.ToResponse())
}

// GetAll handles GET /customers
func (h *customerHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	result, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Customers retrieved successfully", result)
}

// Update handles PUT /customers/:id
func (h *customerHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid customer ID")
	}

	var req UpdateCustomerRequest
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

	return helpers.Success(c, "Customer updated successfully", result.ToResponse())
}

// Delete handles DELETE /customers/:id
func (h *customerHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid customer ID")
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Customer deleted successfully", nil)
}
