package inventory

// CreateInventoryRequest represents the request body for creating inventory
type CreateInventoryRequest struct {
	SKU      string  `json:"sku" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price" validate:"required,min=0"`
}

// UpdateInventoryRequest represents the request body for updating inventory
type UpdateInventoryRequest struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price" validate:"min=0"`
}

// InventoryResponse represents the response for inventory operations
type InventoryResponse struct {
	ID        uint    `json:"id"`
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// ToResponse converts Inventory model to InventoryResponse
func (i *Inventory) ToResponse() InventoryResponse {
	return InventoryResponse{
		ID:        i.ID,
		SKU:       i.SKU,
		Name:      i.Name,
		Quantity:  i.Quantity,
		Price:     i.Price,
		CreatedAt: i.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: i.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// PaginationParams represents pagination request parameters
type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// DefaultPagination returns default pagination params
func DefaultPagination() PaginationParams {
	return PaginationParams{
		Page:     1,
		PageSize: 10,
	}
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
	Data       []InventoryResponse `json:"data"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"page_size"`
	TotalItems int64                `json:"total_items"`
	TotalPages int                  `json:"total_pages"`
}

// ToResponseList converts a slice of Inventory to a slice of InventoryResponse
func ToResponseList(inventories []Inventory) []InventoryResponse {
	result := make([]InventoryResponse, len(inventories))
	for i, inv := range inventories {
		result[i] = inv.ToResponse()
	}
	return result
}
