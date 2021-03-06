package main

import (
	"hewantani/config"
	"hewantani/docs"
	"hewantani/routes"
	"hewantani/services"
	"hewantani/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name   Arief Syaifuddin
// @contact.url    https://github.com/Velezer
// @contact.email  asvelezer@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService  http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "HewanTani API"
	docs.SwaggerInfo.Description = "HewanTani Official API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	Db := config.ConnectDatabase()
	sqlDB, _ := Db.DB()
	defer sqlDB.Close()
	services.Setup(Db)

	r := routes.SetupRouter()
	r.Run()
}
