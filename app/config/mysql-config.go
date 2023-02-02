package config

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// setup database connection
func SetupDatabaseConnection() *gorm.DB {

	databaseConfig, err := GetDatabaseConfig()

	if err != nil {
		panic("Failed to load mysql config")
	}

	DB, err = gorm.Open(mysql.New(mysql.Config{
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
		// log.Fatal("Mysql connect fail...", err)
		panic("Failed to create a connection to mysql")
	}

	return DB
}

// close mysql connection
func CloseMysqlConnection(db *gorm.DB) error {
	dbSql, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database ")
	}

	return dbSql.Close()
}

type DatabaseConfig struct {
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	ConnectionMax int32
}

func GetDatabaseConfig() (DatabaseConfig, error) {

	host := Conf.GetString("database.mysql.host")
	port := Conf.GetString("database.mysql.port")
	database := Conf.GetString("database.mysql.database")
	user := Conf.GetString("database.mysql.user")
	password := Conf.GetString("database.mysql.password")
	connection_max := Conf.GetInt32("database.mysql.connection_max")

	if host == "" || port == "" || database == "" || user == "" || password == "" {
		return DatabaseConfig{
			Host:     host,
			Port:     port,
			Database: database,
			User:     user,
			Password: password,
		}, errors.New("database config missing parameter")
	}

	return DatabaseConfig{
		Host:          host,
		Port:          port,
		Database:      database,
		User:          user,
		Password:      password,
		ConnectionMax: connection_max,
	}, nil
}
