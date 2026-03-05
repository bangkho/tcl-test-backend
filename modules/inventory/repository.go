package inventory

import "gorm.io/gorm"

// InventoryRepository defines the interface for inventory data access
type InventoryRepository interface {
	Create(inventory *Inventory) error
	GetByID(id uint) (*Inventory, error)
	GetAll(page, pageSize int) ([]Inventory, int64, error)
	Update(inventory *Inventory) error
	Delete(id uint) error
	GetBySKU(sku string) (*Inventory, error)
}

// inventoryRepository implements InventoryRepository using GORM
type inventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository creates a new inventory repository
func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

// Create creates a new inventory record
func (r *inventoryRepository) Create(inventory *Inventory) error {
	return r.db.Create(inventory).Error
}

// GetByID retrieves an inventory by ID
func (r *inventoryRepository) GetByID(id uint) (*Inventory, error) {
	var inventory Inventory
	err := r.db.First(&inventory, id).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// GetAll retrieves paginated inventory records
func (r *inventoryRepository) GetAll(page, pageSize int) ([]Inventory, int64, error) {
	var inventories []Inventory
	var total int64

	r.db.Model(&Inventory{}).Count(&total)

	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&inventories).Error

	return inventories, total, err
}

// Update updates an existing inventory record
func (r *inventoryRepository) Update(inventory *Inventory) error {
	return r.db.Save(inventory).Error
}

// Delete deletes an inventory record by ID
func (r *inventoryRepository) Delete(id uint) error {
	return r.db.Delete(&Inventory{}, id).Error
}

// GetBySKU retrieves an inventory by SKU
func (r *inventoryRepository) GetBySKU(sku string) (*Inventory, error) {
	var inventory Inventory
	err := r.db.Where("sku = ?", sku).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}
