package models

import (
	"fmt"
	"go_python_serve/app/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	}

	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database),
	)

	if err != nil {
		log.Fatal("Mysql connect fail...", err)
		panic("Mysql connect fail...")
	} else {
		log.Println("Mysql connect success...")
	}

	// 设置连接池
	Db.DB().SetConnMaxLifetime(100 * time.Second) // 最大连接周期，超时的连接就close
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(20)

	if err = Db.DB().Ping(); err != nil {
		log.Fatal("Mysql connect fail...", err)
		panic("Mysql connect fail...")
	}

	// gorm 根据model 创建表
	gormAutoMigration()
}
