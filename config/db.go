package config

import (
	"hewantani/models"
	"hewantani/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := utils.Getenv("DATABASE_URL", "")
	if utils.Getenv("ENV", "") == "test" {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
