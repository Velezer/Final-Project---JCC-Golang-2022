package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func orderRouter() {
	orderController := controllers.OrderController{}
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.Use(middlewares.JwtAuthMiddleware())
		orderRoutes.Use(middlewares.UserMiddleware())
		orderRoutes.POST("/", orderController.CreateOrder)
	}

	orderMerchantRoutes := r.Group("/orders")
	{
		orderMerchantRoutes.Use(middlewares.JwtAuthMiddleware())
		orderMerchantRoutes.Use(middlewares.SetUserRoleMiddleware())
		orderMerchantRoutes.GET("/", orderController.GetOrders)
	}

	orderOwnerRoutes := r.Group("/orders")
	{
		orderOwnerRoutes.Use(middlewares.JwtAuthMiddleware())
		orderOwnerRoutes.Use(middlewares.OrderOwnerMiddleware())
		orderOwnerRoutes.GET("/:id", orderController.GetOrder)
		orderOwnerRoutes.PUT("/:id/", orderController.UpdateStatusOrder)
		orderOwnerRoutes.DELETE("/:id", orderController.DeleteOrder)
	}
}
