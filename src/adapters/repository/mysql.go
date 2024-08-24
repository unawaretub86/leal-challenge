package repository

import (
	"gorm.io/gorm"
)

type LealRepository struct {
	db *gorm.DB
}

func NewLealRepository(db *gorm.DB) *LealRepository {
	return &LealRepository{
		db: db,
	}
}
