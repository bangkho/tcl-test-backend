package customer

import "gorm.io/gorm"

// CustomerRepository defines the interface for customer data access
type CustomerRepository interface {
	Create(customer *Customer) error
	GetByID(id uint) (*Customer, error)
	GetAll(page, pageSize int) ([]Customer, int64, error)
	Update(customer *Customer) error
	Delete(id uint) error
	GetByEmail(email string) (*Customer, error)
}

// customerRepository implements CustomerRepository using GORM
type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// Create creates a new customer record
func (r *customerRepository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

// GetByID retrieves a customer by ID
func (r *customerRepository) GetByID(id uint) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetAll retrieves paginated customer records
func (r *customerRepository) GetAll(page, pageSize int) ([]Customer, int64, error) {
	var customers []Customer
	var total int64

	r.db.Model(&Customer{}).Count(&total)

	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&customers).Error

	return customers, total, err
}

// Update updates an existing customer record
func (r *customerRepository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

// Delete deletes a customer record by ID
func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&Customer{}, id).Error
}

// GetByEmail retrieves a customer by email
func (r *customerRepository) GetByEmail(email string) (*Customer, error) {
	var customer Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
