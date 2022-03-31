package routes

import "hewantani/controllers"

func userRouter() {
	userController := controllers.UserController{}
	r.POST("/users", userController.Register)
	r.PUT("/users/:id/password", userController.ChangePassword)
	r.POST("/users/login", userController.Login)
}
