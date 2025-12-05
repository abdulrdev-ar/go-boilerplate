package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/inienam06/go-boilerplate/docs" // swag docs
	"github.com/inienam06/go-boilerplate/internal/config"
	"github.com/inienam06/go-boilerplate/internal/core/middleware"
	"github.com/inienam06/go-boilerplate/internal/model"
)

// @title MyApp API
// @version 1.0
// @description Example API with Gin, GORM, Clean Architecture
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and then your token.
func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// For development quick migration (optional)
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("auto migrate failed:", err)
	}

	r := gin.Default()
	r.Use(middleware.ErrorHandling())

	config.RegisterRoutes(r, db)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
