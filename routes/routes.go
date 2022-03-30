package routes

import (
	"hewantani/controllers"
	"hewantani/middlewares"
	"hewantani/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	controller := controllers.Controller{}
	controller.UserService = services.User{Db: db}
	controller.RoleService = services.Role{Db: db}
	controller.StoreService = services.Store{Db: db}
	controller.ProductService = services.Product{Db: db}
	controller.CategoryService = services.Category{Db: db}

	authController := controllers.AuthController{Controller: controller}
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	storeController := controllers.StoreController{Controller: controller}
	storeRoutes := r.Group("/stores")
	storeRoutes.Use(middlewares.JwtAuthMiddleware())
	storeRoutes.Use(middlewares.MerchantMiddleware(controller.UserService))
	storeRoutes.POST("/", storeController.CreateStore)

	productController := controllers.ProductController{Controller: controller}
	productRoutes := r.Group("/products")
	productRoutes.Use(middlewares.JwtAuthMiddleware())
	productRoutes.Use(middlewares.MerchantMiddleware(controller.UserService))
	productRoutes.POST("/", productController.CreateProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
