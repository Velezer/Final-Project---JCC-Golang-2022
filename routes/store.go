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
	{
		merchantRoutes.Use(middlewares.JwtAuthMiddleware())
		merchantRoutes.Use(middlewares.MerchantMiddleware())
		merchantRoutes.POST("/", storeController.CreateStore)
	}

	storeOwnerRoutes := merchantRoutes.Group("/")
	{
		storeOwnerRoutes.Use(middlewares.StoreOwnerMiddleware())
		storeOwnerRoutes.PUT("/:id", storeController.UpdateStore)
		storeOwnerRoutes.DELETE("/:id", storeController.DeleteStore)
	}
}
