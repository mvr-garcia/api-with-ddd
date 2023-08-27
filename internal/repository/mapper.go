package repository

import "acetona/internal/entity"

func MapAssetEntityToModel(entity *entity.Asset) *AssetModel {
	return &AssetModel{
		ID:     entity.ID,
		Name:   entity.Name,
		Symbol: entity.Symbol,
	}
}

func MapAssetModelToEntity(model *AssetModel) *entity.Asset {
	return &entity.Asset{
		ID:     model.ID,
		Name:   model.Name,
		Symbol: model.Symbol,
	}
}
