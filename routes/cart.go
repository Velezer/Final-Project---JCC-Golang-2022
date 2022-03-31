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
	cartRoutes.PATCH("/:id", cartController.AddCartItem)
	cartRoutes.DELETE("/:cart_id/item/:item_id", cartController.DeleteCartItem)
}
