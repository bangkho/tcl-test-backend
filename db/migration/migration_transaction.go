package migration

import (
	"bangkho.dev/tcl/test/backend/modules/transaction"

	"gorm.io/gorm"
)

func MigrationTransaction(db *gorm.DB) error {
	return db.AutoMigrate(&transaction.Transaction{})
}
