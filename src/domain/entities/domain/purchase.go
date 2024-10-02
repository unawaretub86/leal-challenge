package domain

import "time"

type Purchase struct {
	Amount         float64   `json:"amount" validate:"required"`
	UserID         uint64    `json:"user_id" validate:"required"`
	CommerceID     uint64    `json:"commerce_id" validate:"required"`
	BranchID       uint64    `json:"branch_id" validate:"required"`
	CampaignID     uint64    `json:"campaign_id" validate:"required"`
	EarnedPoints   int       `json:"earned_points"`
	EarnedCashBack int       `json:"earned_cashBack"`
	CreatedAt      time.Time `json:"created_at"`
	RedeemPoints   bool      `json:"redeem_points" validate:"required"`
}
