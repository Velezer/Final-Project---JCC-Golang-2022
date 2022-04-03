package config

import (
	"hewantani/models"
	"hewantani/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error
	
	dsn := utils.Getenv("DATABASE_URL", "root:@tcp(127.0.0.1:3306)/db_final_project?charset=utf8mb4&parseTime=True&loc=Local")
	if utils.Getenv("ENV", "") == "production" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

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
