package inventory

import (
	"strconv"

	"bangkho.dev/tcl/test/backend/helpers"

	"github.com/gofiber/fiber/v2"
)

// InventoryHandler handles inventory HTTP requests
type InventoryHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// inventoryHandler implements InventoryHandler
type inventoryHandler struct {
	service InventoryService
}

// NewInventoryHandler creates a new inventory handler
func NewInventoryHandler(service InventoryService) InventoryHandler {
	return &inventoryHandler{service: service}
}

// Create handles POST /inventory
func (h *inventoryHandler) Create(c *fiber.Ctx) error {
	var req CreateInventoryRequest
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

	return helpers.Created(c, "Inventory created successfully", result.ToResponse())
}

// GetByID handles GET /inventory/:id
func (h *inventoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid inventory ID")
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		return helpers.NotFound(c, "Inventory not found")
	}

	return helpers.Success(c, "Inventory retrieved successfully", result.ToResponse())
}

// GetAll handles GET /inventory
func (h *inventoryHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	result, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Inventory retrieved successfully", result)
}

// Update handles PUT /inventory/:id
func (h *inventoryHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid inventory ID")
	}

	var req UpdateInventoryRequest
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

	return helpers.Success(c, "Inventory updated successfully", result.ToResponse())
}

// Delete handles DELETE /inventory/:id
func (h *inventoryHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid inventory ID")
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Inventory deleted successfully", nil)
}
