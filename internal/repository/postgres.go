package repository

import (
	"acetona/internal/entity"

	"gorm.io/gorm"
)

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *assetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset *entity.Asset) (*entity.Asset, error) {
	model := MapAssetEntityToModel(asset)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return MapAssetModelToEntity(model), nil
}
