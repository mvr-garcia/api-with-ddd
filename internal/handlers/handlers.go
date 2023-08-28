package handlers

import (
	"acetona/internal/entity"
	"acetona/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	assetService services.AssetService
}

func NewAPIHandler(assetService services.AssetService) *APIHandler {
	return &APIHandler{assetService: assetService}
}

func (h *APIHandler) CreateAsset(c *gin.Context) {
	var newAsset entity.Asset

	// Extract data from request
	if err := c.ShouldBindJSON(&newAsset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service function
	createdAsset, err := h.assetService.CreateAsset(&newAsset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, createdAsset)
}
