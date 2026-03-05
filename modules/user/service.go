package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your-secret-key-change-in-production")

// UserService defines the interface for user business logic
type UserService interface {
	Register(req *CreateUserRequest) (*User, error)
	Login(req *LoginRequest) (*LoginResponse, error)
	GetByID(id uint) (*User, error)
	GetAll(page, pageSize int) (*PaginationResponse, error)
	Update(id uint, req *UpdateUserRequest) (*User, error)
	Delete(id uint) error
}

// userService implements UserService
type userService struct {
	repository UserRepository
}

// NewUserService creates a new user service
func NewUserService(repository UserRepository) UserService {
	return &userService{repository: repository}
}

// Register creates a new user
func (s *userService) Register(req *CreateUserRequest) (*User, error) {
	// Check if username already exists
	existing, err := s.repository.GetByUsername(req.Username)
	if err == nil && existing != nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set default role
	role := RoleSuperUser
	if req.Role == "admin" {
		role = RoleAdmin
	}

	user := &User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns a JWT token
func (s *userService) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.repository.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
		User:  user.ToResponse(),
	}, nil
}

// GetByID retrieves a user by ID
func (s *userService) GetByID(id uint) (*User, error) {
	return s.repository.GetByID(id)
}

// GetAll retrieves paginated users
func (s *userService) GetAll(page, pageSize int) (*PaginationResponse, error) {
	params := NormalizePagination(PaginationParams{Page: page, PageSize: pageSize})

	users, total, err := s.repository.GetAll(params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / params.PageSize
	if int(total)%params.PageSize > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Data:       ToResponseList(users),
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing user
func (s *userService) Update(id uint, req *UpdateUserRequest) (*User, error) {
	user, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if username is being changed and if it already exists
	if req.Username != "" && req.Username != user.Username {
		existing, err := s.repository.GetByUsername(req.Username)
		if err == nil && existing != nil {
			return nil, errors.New("username already exists")
		}
		user.Username = req.Username
	}

	// Update password if provided
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	// Update role if provided
	if req.Role != "" {
		if req.Role == "admin" {
			user.Role = RoleAdmin
		} else {
			user.Role = RoleSuperUser
		}
	}

	err = s.repository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete deletes a user
func (s *userService) Delete(id uint) error {
	_, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.repository.Delete(id)
}
