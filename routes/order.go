package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func orderRouter() {
	orderController := controllers.OrderController{}
	orderRoutes := r.Group("/orders")
	orderRoutes.Use(middlewares.JwtAuthMiddleware())
	orderRoutes.Use(middlewares.UserMiddleware())
	orderRoutes.POST("/", orderController.CreateOrder)
	orderRoutes.GET("/", orderController.GetOrders)

	orderOwnerRoutes := r.Group("/orders")
	orderRoutes.Use(middlewares.JwtAuthMiddleware())
	orderOwnerRoutes.Use(middlewares.OrderOwnerMiddleware())
	orderOwnerRoutes.PUT("/:id/", orderController.UpdateStatusOrder)
	orderOwnerRoutes.DELETE("/:id", orderController.DeleteOrder)
}
