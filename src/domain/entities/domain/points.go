package domain

import "time"

type Points struct {
	ID               uint       `gorm:"primaryKey"`
	Points           int        `json:"points" validate:"required"`
	CommerceID       uint       `json:"commerce_id" validate:"required"`
	ConversionFactor float64    `json:"conversion_factor" validate:"required"`
	RedeemableItems  string     `json:"redeemable_items" validate:"required"` // Serializado como JSON string
	UserID           uint       `json:"user_id" validate:"required"`
	EarnedDate       time.Time  `json:"earned_date" validate:"required"`
	ExpiryDate       *time.Time `json:"expiry_date,omitempty"`
	CanBeRedeemed    bool       `json:"can_be_redeemed" validate:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}