package internal

import "acetona/internal/entity"

type Repository interface {
	Create(asset *entity.Asset) (*entity.Asset, error)
}
