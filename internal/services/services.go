package services

import (
	"acetona/internal"
	"acetona/internal/entity"
)

type AssetService struct {
	assetRepo internal.Repository
}

func NewAssetService(assetRepo internal.Repository) *AssetService {
	return &AssetService{assetRepo: assetRepo}
}

func (s *AssetService) CreateAsset(asset *entity.Asset) (*entity.Asset, error) {
	return s.assetRepo.Create(asset)
}
