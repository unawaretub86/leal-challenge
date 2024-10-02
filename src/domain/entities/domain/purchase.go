package domain

import "time"

type Purchase struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Amount         float64
	UserID         uint64
	CommerceID     uint64
	BranchID       uint64
	CampaignID     uint64
	EarnedPoints   int
	EarnedCashBack int
	CreatedAt      time.Time
	RedeemPoints   bool
}
