package repository

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

func (r *LealRepository) CreateUser(user domain.User) (*domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *LealRepository) RegisterPurchase(purchase domain.Purchase) (*domain.Purchase, error) {
	if err := r.db.Create(&purchase).Error; err != nil {
		return nil, err
	}

	return &purchase, nil
}
