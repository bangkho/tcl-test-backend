package migration

import (
	"bangkho.dev/tcl/test/backend/modules/user"

	"gorm.io/gorm"
)

func MigrationUser(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{})
}
