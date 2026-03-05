package transaction

import (
	"time"

	"gorm.io/gorm"
)

type TransactionStatus string

type TransactionType string

const (
	TransactionStatusProgress  TransactionStatus = "progress"
	TransactionStatusDone     TransactionStatus = "done"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

const (
	TransactionTypeIn  TransactionType = "in"
	TransactionTypeOut TransactionType = "out"
)

type Transaction struct {
	ID              uint              `gorm:"primaryKey" json:"id"`
	CustomerID      uint              `gorm:"not null;index;foreignKey:CustomerID;references:ID;constraint:OnDelete:CASCADE" json:"customer_id"`
	InventoryID     uint              `gorm:"not null;index;foreignKey:InventoryID;references:ID;constraint:OnDelete:CASCADE" json:"inventory_id"`
	Status          TransactionStatus `gorm:"size:20;default:progress" json:"status"`
	TransactionType TransactionType   `gorm:"size:10;not null" json:"transaction_type"`
	Quantity        int               `gorm:"default:1" json:"quantity"`
	TotalPrice      float64           `gorm:"type:decimal(10,2)" json:"total_price"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"-"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
