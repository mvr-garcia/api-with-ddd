package main

import (
	"acetona/internal/api"
	"acetona/internal/entity"
	"acetona/internal/repository"
	"acetona/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Set up the database connection
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the database tables
	db.AutoMigrate(&entity.Asset{})

	// Create instances for the layers
	assetRepo := repository.NewAssetRepository(db)
	assetService := services.NewAssetService(assetRepo)

	// Set up the Gin router
	router := gin.Default()

	// Set up API routes
	api.SetupRoutes(router, *assetService)

	// Start the server
	router.Run(":8080")
}
