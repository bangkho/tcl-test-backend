package migration

import (
	"bangkho.dev/tcl/test/backend/modules/customer"

	"gorm.io/gorm"
)

func MigrationCustomer(db *gorm.DB) error {
	return db.AutoMigrate(&customer.Customer{})
}
