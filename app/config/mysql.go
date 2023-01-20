package config

import (
	"fmt"
	"gin_server/app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setup database connection
func SetupDatabaseConnection() *gorm.DB {

	var err error

	databaseConfig, err := GetDatabaseConfig()

	if err != nil {
		panic("Mysql connect fail...")
	}

	Db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseConfig.User,
			databaseConfig.Password,
			databaseConfig.Host,
			databaseConfig.Port,
			databaseConfig.Database), // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Mysql connect fail...", err)
		panic("Mysql connect fail...")
	} else {
		log.Println("Mysql connect success...")
	}

	// gorm 根据model 创建表
	Db.AutoMigrate(&models.User{})

	return Db
}

func GetMysqlConfig() {
	panic("unimplemented")
}

// close mysql connection
func CloseMysqlConnection(db *gorm.DB) {
	dbSql, err := db.DB()

	if err != nil {
		panic("Failed close mysql ")
	}

	dbSql.Close()
}
