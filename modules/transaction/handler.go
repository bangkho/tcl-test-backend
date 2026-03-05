package transaction

import (
	"strconv"

	"bangkho.dev/tcl/test/backend/helpers"

	"github.com/gofiber/fiber/v2"
)

// TransactionHandler handles transaction HTTP requests
type TransactionHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// transactionHandler implements TransactionHandler
type transactionHandler struct {
	service TransactionService
}

// NewTransactionHandler creates a new transaction handler
func NewTransactionHandler(service TransactionService) TransactionHandler {
	return &transactionHandler{service: service}
}

// Create handles POST /transactions
func (h *transactionHandler) Create(c *fiber.Ctx) error {
	var req CreateTransactionRequest
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

	return helpers.Created(c, "Transaction created successfully", result.ToResponse())
}

// GetByID handles GET /transactions/:id
func (h *transactionHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid transaction ID")
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		return helpers.NotFound(c, "Transaction not found")
	}

	return helpers.Success(c, "Transaction retrieved successfully", result.ToResponse())
}

// GetAll handles GET /transactions
func (h *transactionHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	result, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Transactions retrieved successfully", result)
}

// Update handles PUT /transactions/:id
func (h *transactionHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid transaction ID")
	}

	var req UpdateTransactionRequest
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

	return helpers.Success(c, "Transaction updated successfully", result.ToResponse())
}

// Delete handles DELETE /transactions/:id
func (h *transactionHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.BadRequest(c, "Invalid transaction ID")
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, "Transaction deleted successfully", nil)
}
