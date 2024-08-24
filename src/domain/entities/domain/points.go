package domain

import "time"

type Points struct {
	ID               uint64     `gorm:"primaryKey"`
	Points           int        `json:"points" validate:"required"`
	CommerceID       uint64     `json:"commerce_id" validate:"required"`
	ConversionFactor float64    `json:"conversion_factor" validate:"required"`
	RedeemableItems  uint64     `json:"redeemable_items" validate:"required"`
	UserID           uint64     `json:"user_id" validate:"required" gorm:"index"`
	EarnedDate       time.Time  `json:"earned_date" validate:"required"`
	ExpiryDate       *time.Time `json:"expiry_date,omitempty"`
	CanBeRedeemed    bool       `json:"can_be_redeemed" validate:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
