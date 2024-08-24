package dto

import (
	"errors"
	"time"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type (
	CreateCommerceDTO struct {
		Name string `json:"name" validate:"required"`
	}

	CreateBranchDTO struct {
		Name       string    `json:"name" validate:"required"`
		CommerceID uint      `json:"commerce_id" validate:"required"`
		Address    string    `json:"address" validate:"required"`
		ActiveDate time.Time `json:"active_date" validate:"required"`
	}

	CreateCampaignDTO struct {
		Name        string    `json:"name" validate:"required"`
		CommerceID  uint      `json:"commerce_id" validate:"required"`
		BranchID    *uint     `json:"branch_id"`
		StartDate   time.Time `json:"start_date" validate:"required"`
		EndDate     time.Time `json:"end_date" validate:"required"`
		IsForAll    bool      `json:"is_for_all"`
		BonusFactor float64   `json:"bonus_factor" validate:"required"`
		BonusType   string    `json:"bonus_type" validate:"required"`
		MinPurchase float64   `json:"min_purchase" validate:"required"`
	}
)

func NewCommerce(commerce CreateCommerceDTO) *domain.Commerce {
	return &domain.Commerce{
		Name: commerce.Name,
	}
}

func NewBranch(dto CreateBranchDTO) *domain.Branch {
	return &domain.Branch{
		Name:       dto.Name,
		CommerceID: dto.CommerceID,
		Address:    dto.Address,
		StartDate:  dto.ActiveDate,
	}
}

func NewCampaign(campaign CreateCampaignDTO) (*domain.Campaign, error) {
	if campaign.EndDate.Before(time.Now()) {
		return nil, errors.New("end_date cannot be in the past")
	}

	return &domain.Campaign{
		Name:        campaign.Name,
		CommerceID:  campaign.CommerceID,
		BranchID:    campaign.BranchID,
		BonusFactor: campaign.BonusFactor,
		BonusType:   campaign.BonusType,
		MinPurchase: &campaign.MinPurchase,
		StartDate:   campaign.StartDate,
		EndDate:     campaign.EndDate,
		IsForAll:    campaign.IsForAll,
	}, nil
}
