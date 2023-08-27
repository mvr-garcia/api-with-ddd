package handlers

import "acetona/internal/services"

type APIHandler struct {
	assetService services.AssetService
}

func NewAPIHandler(assetService services.AssetService) *APIHandler {
	return &APIHandler{assetService: assetService}
}
