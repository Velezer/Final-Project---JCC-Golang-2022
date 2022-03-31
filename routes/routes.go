package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var r *gin.Engine = gin.Default()


func cartRouter() {
	cartController := controllers.CartController{}
	cartRoutes := r.Group("/carts")
	cartRoutes.Use(middlewares.JwtAuthMiddleware())
	cartRoutes.Use(middlewares.UserMiddleware())
	cartRoutes.POST("/", cartController.CreateCart)
	cartRoutes.PATCH("/", cartController.AddCartItem)
}



func SetupRouter() *gin.Engine {
	userRouter()
	storeRouter()
	productRouter()
	cartRouter()
	orderRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
