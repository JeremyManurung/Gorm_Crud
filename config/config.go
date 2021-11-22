package config

import (
	"res-task/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_echo?charset=utf8mb4&parseTime=True&loc=Local"
  	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}
}
func InitMigration() {
	DB.AutoMigrate(
		&model.User{},
	)
}