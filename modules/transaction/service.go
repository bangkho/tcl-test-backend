package transaction

import (
	"errors"

	customerRepo "bangkho.dev/tcl/test/backend/modules/customer"
	inventoryRepo "bangkho.dev/tcl/test/backend/modules/inventory"

	"gorm.io/gorm"
)

// TransactionService defines the interface for transaction business logic
type TransactionService interface {
	Create(req *CreateTransactionRequest) (*Transaction, error)
	GetByID(id uint) (*Transaction, error)
	GetAll(page, pageSize int) (*PaginationResponse, error)
	Update(id uint, req *UpdateTransactionRequest) (*Transaction, error)
	Delete(id uint) error
}

// transactionService implements TransactionService
type transactionService struct {
	repository      TransactionRepository
	customerService customerRepo.CustomerRepository
	inventoryService inventoryRepo.InventoryRepository
}

// NewTransactionService creates a new transaction service
func NewTransactionService(
	repository TransactionRepository,
	customerService customerRepo.CustomerRepository,
	inventoryService inventoryRepo.InventoryRepository,
) TransactionService {
	return &transactionService{
		repository:      repository,
		customerService: customerService,
		inventoryService: inventoryService,
	}
}

// Create creates a new transaction
func (s *transactionService) Create(req *CreateTransactionRequest) (*Transaction, error) {
	// Validate customer exists
	_, err := s.customerService.GetByID(req.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// Validate inventory exists
	inventory, err := s.inventoryService.GetByID(req.InventoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, err
	}

	// Check inventory quantity for "out" transactions
	if req.TransactionType == TransactionTypeOut {
		if inventory.Quantity < req.Quantity {
			return nil, errors.New("insufficient inventory quantity")
		}
	}

	// Calculate total price
	totalPrice := inventory.Price * float64(req.Quantity)

	transaction := &Transaction{
		CustomerID:      req.CustomerID,
		InventoryID:     req.InventoryID,
		TransactionType: req.TransactionType,
		Quantity:        req.Quantity,
		TotalPrice:      totalPrice,
		Status:          TransactionStatusProgress,
	}

	err = s.repository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// GetByID retrieves a transaction by ID
func (s *transactionService) GetByID(id uint) (*Transaction, error) {
	return s.repository.GetByID(id)
}

// GetAll retrieves paginated transactions
func (s *transactionService) GetAll(page, pageSize int) (*PaginationResponse, error) {
	params := NormalizePagination(PaginationParams{Page: page, PageSize: pageSize})

	transactions, total, err := s.repository.GetAll(params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / params.PageSize
	if int(total)%params.PageSize > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Data:       ToResponseList(transactions),
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing transaction
func (s *transactionService) Update(id uint, req *UpdateTransactionRequest) (*Transaction, error) {
	transaction, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}

	if req.Status != "" {
		transaction.Status = req.Status
	}

	if req.TransactionType != "" {
		transaction.TransactionType = req.TransactionType
	}

	if req.Quantity > 0 {
		transaction.Quantity = req.Quantity
	}

	err = s.repository.Update(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Delete deletes a transaction
func (s *transactionService) Delete(id uint) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("transaction not found")
		}
		return err
	}

	return s.repository.Delete(id)
}
