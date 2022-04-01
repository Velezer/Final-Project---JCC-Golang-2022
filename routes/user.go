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
	afterLoginRoutes.GET("/", userController.GetUser)                // get user based on jwt
	afterLoginRoutes.PUT("/", userController.UpdateUser)             // edit user info, but you can't change the role
	afterLoginRoutes.PUT("/password", userController.ChangePassword) // change password only
}
