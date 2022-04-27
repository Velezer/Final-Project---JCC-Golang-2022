package tests

import (
	"hewantani/config"
	"hewantani/utils"

	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	utils.SetEnv("ENV", "test")
	return config.ConnectDatabase()
}
