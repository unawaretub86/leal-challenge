package domain

import "time"

type Campaign struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required"`
	CommerceID  uint      `json:"commerce_id" validate:"required"`
	BranchID    uint      `json:"branch_id" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	BonusType   string    `json:"bonus_type"`
	BonusFactor float64   `json:"bonus_factor"`
	MinPurchase *float64  `json:"min_purchase,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
