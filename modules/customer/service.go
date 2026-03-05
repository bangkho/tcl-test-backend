package customer

import (
	"errors"

	"gorm.io/gorm"
)

// CustomerService defines the interface for customer business logic
type CustomerService interface {
	Create(req *CreateCustomerRequest) (*Customer, error)
	GetByID(id uint) (*Customer, error)
	GetAll(page, pageSize int) (*PaginationResponse, error)
	Update(id uint, req *UpdateCustomerRequest) (*Customer, error)
	Delete(id uint) error
}

// customerService implements CustomerService
type customerService struct {
	repository CustomerRepository
}

// NewCustomerService creates a new customer service
func NewCustomerService(repository CustomerRepository) CustomerService {
	return &customerService{repository: repository}
}

// Create creates a new customer
func (s *customerService) Create(req *CreateCustomerRequest) (*Customer, error) {
	// Check if email already exists
	if req.Email != "" {
		existing, err := s.repository.GetByEmail(req.Email)
		if err == nil && existing != nil {
			return nil, errors.New("email already exists")
		}
	}

	customer := &Customer{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	err := s.repository.Create(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// GetByID retrieves a customer by ID
func (s *customerService) GetByID(id uint) (*Customer, error) {
	return s.repository.GetByID(id)
}

// GetAll retrieves paginated customers
func (s *customerService) GetAll(page, pageSize int) (*PaginationResponse, error) {
	params := NormalizePagination(PaginationParams{Page: page, PageSize: pageSize})

	customers, total, err := s.repository.GetAll(params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / params.PageSize
	if int(total)%params.PageSize > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Data:       ToResponseList(customers),
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing customer
func (s *customerService) Update(id uint, req *UpdateCustomerRequest) (*Customer, error) {
	customer, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// Check if email is being changed and if it already exists
	if req.Email != "" && req.Email != customer.Email {
		existing, err := s.repository.GetByEmail(req.Email)
		if err == nil && existing != nil {
			return nil, errors.New("email already exists")
		}
		customer.Email = req.Email
	}

	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Address != "" {
		customer.Address = req.Address
	}

	err = s.repository.Update(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Delete deletes a customer
func (s *customerService) Delete(id uint) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		return err
	}

	return s.repository.Delete(id)
}
