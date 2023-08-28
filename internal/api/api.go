package api

import (
	"acetona/internal/handlers"
	"acetona/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, assetService services.AssetService) {

	handler := handlers.NewAPIHandler(assetService)

	// Set up API routes
	router.POST("/assets", handler.CreateAsset)

}
