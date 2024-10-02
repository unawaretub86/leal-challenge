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
		CommerceID uint64    `json:"commerce_id" validate:"required"`
		Address    string    `json:"address" validate:"required"`
		ActiveDate time.Time `json:"active_date" validate:"required"`
	}

	CreateCampaignDTO struct {
		Name        string                 `json:"name" validate:"required"`
		CommerceID  uint64                 `json:"commerce_id" validate:"required"`
		BranchID    uint64                 `json:"branch_id"`
		StartDate   time.Time              `json:"start_date" validate:"required"`
		EndDate     time.Time              `json:"end_date" validate:"required"`
		IsForAll    bool                   `json:"is_for_all"`
		Awards      []CreateAwardsCampaign `json:"awards,omitempty"`
		BonusFactor float64                `json:"bonus_factor" validate:"required"`
		BonusType   string                 `json:"bonus_type" validate:"required,oneof=double percentage"`
		MinPurchase float64                `json:"min_purchase" validate:"required"`
	}

	CreateAwardsCampaign struct {
		ID          uint64  `json:"id"`
		CommerceID  uint64  `json:"commerce_id"`
		Description string  `json:"description"`
		PointsCost  uint64  `json:"points_cost"`
		Value       float64 `json:"value"`
		CampaignID  uint64
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
		BranchID:    &campaign.BranchID,
		BonusFactor: campaign.BonusFactor,
		BonusType:   campaign.BonusType,
		MinPurchase: &campaign.MinPurchase,
		StartDate:   campaign.StartDate,
		EndDate:     campaign.EndDate,
		IsForAll:    campaign.IsForAll,
		Award:       getAwards(campaign.Awards),
	}, nil
}

func getAwards(awardsDTO []CreateAwardsCampaign) []domain.Award {
	var awards []domain.Award

	for _, awardDTO := range awardsDTO {
		award := domain.Award{
			CommerceID:  awardDTO.CommerceID,
			Description: awardDTO.Description,
			PointsCost:  awardDTO.PointsCost,
			Value:       awardDTO.Value,
		}
		awards = append(awards, award)
	}

	return awards
}
