package database

import (
	"fmt"

	"clean_architecture/config"
	"clean_architecture/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ca *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ca.DBUser,
		ca.DBPwd,
		ca.DBHost,
		ca.DBPort,
		ca.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error:", err.Error())
		return nil
	}
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}
