package repository

type AssetModel struct {
	ID     uint `gorm:"primary_key"`
	Name   string
	Symbol string
}
