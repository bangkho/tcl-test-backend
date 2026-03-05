package migration

import (
	"bangkho.dev/tcl/test/backend/modules/inventory"

	"gorm.io/gorm"
)

func MigrationInventory(db *gorm.DB) error {
	return db.AutoMigrate(&inventory.Inventory{})
}
