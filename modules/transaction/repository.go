package transaction

import "gorm.io/gorm"

// TransactionRepository defines the interface for transaction data access
type TransactionRepository interface {
	Create(transaction *Transaction) error
	GetByID(id uint) (*Transaction, error)
	GetAll(page, pageSize int) ([]Transaction, int64, error)
	Update(transaction *Transaction) error
	Delete(id uint) error
}

// transactionRepository implements TransactionRepository using GORM
type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new transaction repository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// Create creates a new transaction record
func (r *transactionRepository) Create(transaction *Transaction) error {
	return r.db.Create(transaction).Error
}

// GetByID retrieves a transaction by ID
func (r *transactionRepository) GetByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetAll retrieves paginated transaction records
func (r *transactionRepository) GetAll(page, pageSize int) ([]Transaction, int64, error) {
	var transactions []Transaction
	var total int64

	r.db.Model(&Transaction{}).Count(&total)

	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&transactions).Error

	return transactions, total, err
}

// Update updates an existing transaction record
func (r *transactionRepository) Update(transaction *Transaction) error {
	return r.db.Save(transaction).Error
}

// Delete deletes a transaction record by ID
func (r *transactionRepository) Delete(id uint) error {
	return r.db.Delete(&Transaction{}, id).Error
}
