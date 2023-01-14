package models

// gormAutoMigration gotm autoMigrate
func gormAutoMigration() {
	Db.AutoMigrate(&User{})
}
