package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func productRouter() {
	productController := controllers.ProductController{}

	productRoutes := r.Group("/products")
	productRoutes.GET("/", productController.GetProducts)
	productRoutes.GET("/:id", productController.GetProduct)

	merchantRoutes := productRoutes.Group("/")
	merchantRoutes.Use(middlewares.JwtAuthMiddleware())
	merchantRoutes.Use(middlewares.MerchantMiddleware())
	merchantRoutes.POST("/", productController.CreateProduct)

	productOwnerRoutes := merchantRoutes.Group("/")
	productOwnerRoutes.Use(middlewares.ProductOwnerMiddleware())
	productOwnerRoutes.PUT("/:id", productController.UpdateProduct)
	productOwnerRoutes.DELETE("/:id", productController.DeleteProduct)

}
