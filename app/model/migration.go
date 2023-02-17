package model

import "gorm.io/gorm"

// gormAutoMigration gorm autoMigrate
func GormAutoMigration(Db *gorm.DB) {
	// auto migrate
	Db.AutoMigrate(
		&Role{},
		&User{},
		// &LoginRecord{},
		&Todo{},
	)
}
