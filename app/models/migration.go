package models

import "gorm.io/gorm"

// gormAutoMigration gorm autoMigrate
func GormAutoMigration(Db *gorm.DB) {
	// auto migrate
	Db.AutoMigrate(
		&User{},
		&LoginRecord{},
		&Todo{},
	)
}
