package domain

import "time"

type (
	Campaign struct {
		ID          uint64 `gorm:"primaryKey"`
		Name        string
		CommerceID  uint64
		BranchID    *uint64
		StartDate   time.Time
		EndDate     time.Time
		IsForAll    bool
		BonusType   string
		BonusFactor float64
		MinPurchase *float64
		Award       []Award `gorm:"foreignKey:CampaignID"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Campaigns []Campaign
)
