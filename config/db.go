package config

import (
	"fmt"
	"hewantani/models"
	"hewantani/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("DB_USERNAME", "root")
	password := utils.Getenv("DB_PASSWORD", "")
	host := utils.Getenv("DB_HOST", "tcp(127.0.0.1:3306)")
	database := "db_final_project"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Cart{},
		&models.CartItem{},
		&models.Category{},
		&models.Product{},
		&models.Store{},
		&models.OrderStatus{},
		&models.Order{},
	)

	if db.Migrator().HasTable(&models.Role{}) {
		db.Create(&models.Role{Name: models.ROLE_USER})
		db.Create(&models.Role{Name: models.ROLE_MERCHANT})
	}
	if db.Migrator().HasTable(&models.OrderStatus{}) {
		db.Create(&models.OrderStatus{Name: models.ORDER_UNPAID})
		db.Create(&models.OrderStatus{Name: models.ORDER_PAID})
		db.Create(&models.OrderStatus{Name: models.ORDER_CANCELLED})
		db.Create(&models.OrderStatus{Name: models.ORDER_SHIPPING})
		db.Create(&models.OrderStatus{Name: models.ORDER_DELIVERED})
	}

	return db
}
