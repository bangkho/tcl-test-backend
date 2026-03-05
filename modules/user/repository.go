package user

import "gorm.io/gorm"

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetAll(page, pageSize int) ([]User, int64, error)
	Update(user *User) error
	Delete(id uint) error
	GetByUsername(username string) (*User, error)
}

// userRepository implements UserRepository using GORM
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user record
func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves paginated user records
func (r *userRepository) GetAll(page, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	r.db.Model(&User{}).Count(&total)

	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error

	return users, total, err
}

// Update updates an existing user record
func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user record by ID
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

// GetByUsername retrieves a user by username
func (r *userRepository) GetByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
