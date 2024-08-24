package domain

import "time"

type Cashback struct {
	ID               uint64     `gorm:"primaryKey"`
	Coins            int        `json:"coins" validate:"required"`
	CommerceID       uint64     `json:"commerce_id" validate:"required"`
	ConversionFactor float64    `json:"conversion_factor" validate:"required"`
	Redeemable       bool       `json:"redeemable" validate:"required"`
	UserID           uint64     `json:"user_id" validate:"required"`
	EarnedDate       time.Time  `json:"earned_date" validate:"required"`
	ExpiryDate       *time.Time `json:"expiry_date,omitempty"`
	EquivalentValue  float64    `json:"equivalent_value" validate:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
