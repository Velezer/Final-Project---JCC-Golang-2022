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
	cartRoutes.GET("/", cartController.GetUserCarts)
	cartRoutes.POST("/", cartController.CreateCart)

	cartOwnerRoutes := cartRoutes.Group("/")
	cartOwnerRoutes.Use(middlewares.CartOwnerMiddleware())
	cartOwnerRoutes.GET("/:id", cartController.GetUserCart)
	cartOwnerRoutes.PUT("/:id", cartController.UpdateCart)
	cartOwnerRoutes.DELETE("/:id", cartController.DeleteCart)
	cartOwnerRoutes.PUT("/:id/items", cartController.UpdateCartItem)
}
