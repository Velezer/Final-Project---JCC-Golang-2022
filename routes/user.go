package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func userRouter() {
	userController := controllers.UserController{}
	r.POST("/users", userController.Register)
	r.POST("/users/login", userController.Login)

	afterLoginRoutes := r.Group("/users")
	afterLoginRoutes.Use(middlewares.JwtAuthMiddleware())
	afterLoginRoutes.PUT("/users", userController.UpdateUser)              // edit user info, but you can't change the role
	afterLoginRoutes.PUT("/users/password", userController.ChangePassword) // change password only
}
