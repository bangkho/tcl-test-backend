package transaction

// CreateTransactionRequest represents the request body for creating transaction
type CreateTransactionRequest struct {
	CustomerID      uint            `json:"customer_id" validate:"required"`
	InventoryID     uint            `json:"inventory_id" validate:"required"`
	TransactionType TransactionType `json:"transaction_type" validate:"required,oneof=in out"`
	Quantity        int             `json:"quantity" validate:"required,min=1"`
}

// UpdateTransactionRequest represents the request body for updating transaction
type UpdateTransactionRequest struct {
	Status          TransactionStatus `json:"status"`
	TransactionType TransactionType   `json:"transaction_type"`
	Quantity        int               `json:"quantity"`
}

// TransactionResponse represents the response for transaction operations
type TransactionResponse struct {
	ID              uint    `json:"id"`
	CustomerID      uint    `json:"customer_id"`
	InventoryID     uint    `json:"inventory_id"`
	Status          string  `json:"status"`
	TransactionType string  `json:"transaction_type"`
	Quantity        int     `json:"quantity"`
	TotalPrice      float64 `json:"total_price"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// ToResponse converts Transaction model to TransactionResponse
func (t *Transaction) ToResponse() TransactionResponse {
	return TransactionResponse{
		ID:              t.ID,
		CustomerID:      t.CustomerID,
		InventoryID:     t.InventoryID,
		Status:          string(t.Status),
		TransactionType: string(t.TransactionType),
		Quantity:        t.Quantity,
		TotalPrice:      t.TotalPrice,
		CreatedAt:       t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       t.UpdatedAt.Format("2006-01-02 15:04:05"),
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
	Data       []TransactionResponse `json:"data"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	TotalItems int64                 `json:"total_items"`
	TotalPages int                   `json:"total_pages"`
}

// ToResponseList converts a slice of Transaction to a slice of TransactionResponse
func ToResponseList(transactions []Transaction) []TransactionResponse {
	result := make([]TransactionResponse, len(transactions))
	for i, t := range transactions {
		result[i] = t.ToResponse()
	}
	return result
}
