package config

import (
	"github.com/gin-gonic/gin"
	"github.com/inienam06/go-boilerplate/internal/modules/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// RegisterRoutes registers controllers to gin.Engine
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := user.InitUserRepository(db)
	userService := user.InitUserService(userRepo)
	userController := user.InitUserController(userService)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/users")
	{
		users.POST("", userController.CreateUser)
		users.GET("", userController.ListUsers)
		users.GET("/:id", userController.GetUser)
	}
}
