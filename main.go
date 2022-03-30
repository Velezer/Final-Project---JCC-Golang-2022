package main

import (
	"hewantani/config"
	"hewantani/docs"
	"hewantani/routes"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService  http://swagger.io/terms/

func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "HewanTani API"
	docs.SwaggerInfo.Description = "HewanTani Official API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
