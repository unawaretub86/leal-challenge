package repository

import (

	"gorm.io/gorm"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type LealRepository struct {
	db *gorm.DB
}

func NewLealRepository(db *gorm.DB) *LealRepository {
	return &LealRepository{
		db: db,
	}
}

func (r *LealRepository) CreateUser(user domain.User) (*domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
