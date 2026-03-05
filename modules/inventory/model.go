package inventory

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	SKU       string        `gorm:"uniqueIndex;size:50;not null" json:"sku"`
	Name      string        `gorm:"size:100;not null" json:"name"`
	Quantity  int           `gorm:"default:0" json:"quantity"`
	Price     float64       `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (i *Inventory) TableName() string {
	return "inventory"
}
