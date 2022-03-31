package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func storeRouter() {
	storeController := controllers.StoreController{}
	storeRoutes := r.Group("/stores")
	storeRoutes.GET("/", storeController.GetStores)

	merchantRoutes := storeRoutes.Group("/")
	merchantRoutes.Use(middlewares.JwtAuthMiddleware())
	merchantRoutes.Use(middlewares.MerchantMiddleware())
	merchantRoutes.POST("/", storeController.CreateStore)
	merchantRoutes.PUT("/:id", storeController.UpdateStore)
	merchantRoutes.DELETE("/:id", storeController.DeleteStore)
}
