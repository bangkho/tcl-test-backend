package inventory

import (
	"errors"

	"gorm.io/gorm"
)

// InventoryService defines the interface for inventory business logic
type InventoryService interface {
	Create(req *CreateInventoryRequest) (*Inventory, error)
	GetByID(id uint) (*Inventory, error)
	GetAll(page, pageSize int) (*PaginationResponse, error)
	Update(id uint, req *UpdateInventoryRequest) (*Inventory, error)
	Delete(id uint) error
}

// inventoryService implements InventoryService
type inventoryService struct {
	repository InventoryRepository
}

// NewInventoryService creates a new inventory service
func NewInventoryService(repository InventoryRepository) InventoryService {
	return &inventoryService{repository: repository}
}

// Create creates a new inventory item
func (s *inventoryService) Create(req *CreateInventoryRequest) (*Inventory, error) {
	// Check if SKU already exists
	existing, err := s.repository.GetBySKU(req.SKU)
	if err == nil && existing != nil {
		return nil, errors.New("SKU already exists")
	}

	inventory := &Inventory{
		SKU:      req.SKU,
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	err = s.repository.Create(inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// GetByID retrieves an inventory item by ID
func (s *inventoryService) GetByID(id uint) (*Inventory, error) {
	return s.repository.GetByID(id)
}

// GetAll retrieves paginated inventory items
func (s *inventoryService) GetAll(page, pageSize int) (*PaginationResponse, error) {
	// Normalize pagination params
	params := NormalizePagination(PaginationParams{Page: page, PageSize: pageSize})

	inventories, total, err := s.repository.GetAll(params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / params.PageSize
	if int(total)%params.PageSize > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Data:       ToResponseList(inventories),
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing inventory item
func (s *inventoryService) Update(id uint, req *UpdateInventoryRequest) (*Inventory, error) {
	inventory, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, err
	}

	// Check if SKU is being changed and if it already exists
	if req.SKU != "" && req.SKU != inventory.SKU {
		existing, err := s.repository.GetBySKU(req.SKU)
		if err == nil && existing != nil {
			return nil, errors.New("SKU already exists")
		}
		inventory.SKU = req.SKU
	}

	if req.Name != "" {
		inventory.Name = req.Name
	}
	if req.Quantity != 0 {
		inventory.Quantity = req.Quantity
	}
	if req.Price != 0 {
		inventory.Price = req.Price
	}

	err = s.repository.Update(inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// Delete deletes an inventory item
func (s *inventoryService) Delete(id uint) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("inventory not found")
		}
		return err
	}

	return s.repository.Delete(id)
}
