package config

import (
	"fmt"
	"hewantani/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	database := "db_final_project"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&models.Cart{},
		&models.Category{},
		&models.Product{},
		&models.Store{},
		&models.User{},
		&models.Order{},
		&models.Role{},
	)

	if db.Migrator().HasTable(&models.Role{}) {
		db.Create(&models.Role{Name: "USER"})
		db.Create(&models.Role{Name: "MERCHANT"})
	}

	return db
}
