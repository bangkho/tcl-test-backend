package customer


// CreateCustomerRequest represents the request body for creating customer
type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// UpdateCustomerRequest represents the request body for updating customer
type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email" validate:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// CustomerResponse represents the response for customer operations
type CustomerResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToResponse converts Customer model to CustomerResponse
func (c *Customer) ToResponse() CustomerResponse {
	return CustomerResponse{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   c.Address,
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: c.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// PaginationParams represents pagination request parameters
type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// NormalizePagination ensures pagination params are within valid bounds
func NormalizePagination(params PaginationParams) PaginationParams {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 {
		params.PageSize = 10
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}
	return params
}

// PaginationResponse represents pagination response
type PaginationResponse struct {
	Data       []CustomerResponse `json:"data"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
	TotalItems int64              `json:"total_items"`
	TotalPages int                `json:"total_pages"`
}

// ToResponseList converts a slice of Customer to a slice of CustomerResponse
func ToResponseList(customers []Customer) []CustomerResponse {
	result := make([]CustomerResponse, len(customers))
	for i, c := range customers {
		result[i] = c.ToResponse()
	}
	return result
}
