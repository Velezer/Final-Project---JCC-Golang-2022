package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
)

func userRouter() {
	userController := controllers.UserController{}
	r.POST("/user", userController.Register)
	r.POST("/user/login", userController.Login)

	afterLoginRoutes := r.Group("/user")
	{
		afterLoginRoutes.Use(middlewares.JwtAuthMiddleware())
		afterLoginRoutes.GET("/", userController.GetUser)    // get user based on jwt
		afterLoginRoutes.PUT("/", userController.UpdateUser) // edit user info, but you can't change the role
		afterLoginRoutes.DELETE("/", userController.DeleteUser)
		afterLoginRoutes.PUT("/password", userController.ChangePassword) // change password only
	}
}
