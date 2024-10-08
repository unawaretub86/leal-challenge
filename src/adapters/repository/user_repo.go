package repository

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"gorm.io/gorm"
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

	user, err := r.GetUser(purchase.UserID)
	if err != nil {
		return nil, err
	}

	user.Cashback += purchase.EarnedCashBack
	user.Points += purchase.EarnedPoints

	_, err = r.UpdateUser(purchase.UserID, *user)
	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r *LealRepository) GetUser(userID uint64) (*domain.User, error) {
	resultData, err := r.getUser(userID)
	if err.Error != nil {
		return nil, err.Error
	}

	return resultData, nil
}

func (r *LealRepository) UpdateUser(userID uint64, user domain.User) (*domain.User, error) {
	resultData, err := r.getUser(userID)
	if err.Error != nil {
		return nil, err.Error
	}

	user.ID = resultData.ID

	result := r.db.Where("id =?", user.ID).
		Updates(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *LealRepository) GetPurchase(purchaseID uint64) (*domain.Purchase, error) {
	purchase := &domain.Purchase{}

	result := r.db.
		Where("id = ?", purchaseID).
		Take(&purchase)

	if result.Error != nil {
		return nil, result.Error
	}

	return purchase, nil
}

func (r *LealRepository) Redeem(redeem domain.Redeem) (*domain.Redeem, error) {
	if err := r.db.Create(&redeem).Error; err != nil {
		return nil, err
	}

	user, err := r.GetUser(redeem.UserID)
	if err != nil {
		return nil, err
	}

	user.Cashback -= int(redeem.RedeemedCashBack)
	user.Points -= int(redeem.RedeemedPoints)

	_, err = r.UpdateUser(redeem.UserID, *user)
	if err != nil {
		return nil, err
	}

	return &redeem, nil
}

func (r *LealRepository) getUser(id uint64) (*domain.User, *gorm.DB) {
	user := &domain.User{}

	result := r.db.
		Where("id = ?", id).
		Take(&user)

	return user, result
}
