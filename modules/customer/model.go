package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Name      string        `gorm:"size:100;not null" json:"name"`
	Email     string        `gorm:"uniqueIndex;size:100" json:"email"`
	Phone     string        `gorm:"size:20" json:"phone"`
	Address   string        `gorm:"type:text" json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Customer) TableName() string {
	return "customer"
}
