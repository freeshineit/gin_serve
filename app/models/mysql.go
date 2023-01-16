package models

import (
	"fmt"
	"go_python_serve/app/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db gorm instance
var Db *gorm.DB

type ConfigMysql struct {
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	ConnectionMax int32
}

func GetMysqlConfig() (ConfigMysql, error) {

	return ConfigMysql{
		Host:          config.Conf.GetString("database.mysql.host"),
		Port:          config.Conf.GetString("database.mysql.port"),
		Database:      config.Conf.GetString("database.mysql.database"),
		User:          config.Conf.GetString("database.mysql.user"),
		Password:      config.Conf.GetString("database.mysql.password"),
		ConnectionMax: config.Conf.GetInt32("database.mysql.connection_max"),
	}, nil
}

// InitMysqlDB  init mysql db
func InitMysqlDB() {

	var err error

	mysqlConfig, err := GetMysqlConfig()

	if err != nil {
		panic("Mysql connect fail...")
	}

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlConfig.User,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Database), // data source name
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

	// if err = Db.DB().Ping(); err != nil {
	// 	log.Fatal("Mysql connect fail...", err)
	// 	panic("Mysql connect fail...")
	// }

	// gorm 根据model 创建表
	gormAutoMigration()
}
