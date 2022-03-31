package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var r *gin.Engine = gin.Default()



func storeRouter() {
	storeController := controllers.StoreController{}
	storeRoutes := r.Group("/stores")
	storeRoutes.Use(middlewares.JwtAuthMiddleware())
	storeRoutes.Use(middlewares.MerchantMiddleware())
	storeRoutes.POST("/", storeController.CreateStore)
}

func prouctRouter() {
	productController := controllers.ProductController{}
	productRoutes := r.Group("/products")
	productRoutes.Use(middlewares.JwtAuthMiddleware())
	productRoutes.Use(middlewares.MerchantMiddleware())
	productRoutes.POST("/", productController.CreateProduct)
}

func cartRouter() {
	cartController := controllers.CartController{}
	cartRoutes := r.Group("/carts")
	cartRoutes.Use(middlewares.JwtAuthMiddleware())
	cartRoutes.Use(middlewares.UserMiddleware())
	cartRoutes.POST("/", cartController.CreateCart)
	cartRoutes.PATCH("/", cartController.AddCartItem)
}

func orderRouter() {
	orderController := controllers.OrderController{}
	orderRoutes := r.Group("/orders")
	orderRoutes.Use(middlewares.JwtAuthMiddleware())
	orderRoutes.Use(middlewares.UserMiddleware())
	orderRoutes.POST("/", orderController.CreateOrder)
}

func SetupRouter() *gin.Engine {
	userRouter()
	storeRouter()
	prouctRouter()
	cartRouter()
	orderRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
