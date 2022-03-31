package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func productRouter() {
	productController := controllers.ProductController{}

	productRoutes := r.Group("/products")
	productRoutes.GET("/", productController.GetProducts)

	merchantRoutes := productRoutes.Group("/")
	merchantRoutes.Use(middlewares.JwtAuthMiddleware())
	merchantRoutes.Use(middlewares.MerchantMiddleware())
	merchantRoutes.POST("/", productController.CreateProduct)

}
