package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func cartRouter() {
	cartController := controllers.CartController{}
	cartRoutes := r.Group("/carts")
	cartRoutes.Use(middlewares.JwtAuthMiddleware())
	cartRoutes.Use(middlewares.UserMiddleware())
	cartRoutes.GET("/", cartController.GetUserCart)
	cartRoutes.POST("/", cartController.CreateCart)

	cartOwnerRoutes := cartRoutes.Group("/")
	cartOwnerRoutes.Use(middlewares.CartOwnerMiddleware())
	cartOwnerRoutes.PUT("/:id", cartController.UpdateCart)
	cartOwnerRoutes.POST("/:id/items", cartController.AddCartItem)
	// cartRoutes.DELETE("/:cart_id/item/:item_id", cartController.DeleteCartItem)
}
