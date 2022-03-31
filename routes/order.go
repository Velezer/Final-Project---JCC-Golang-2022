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
	orderRoutes.DELETE("/:id", orderController.DeleteOrder)
}