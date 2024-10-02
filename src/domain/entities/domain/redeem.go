package domain

type Redeem struct {
	UserID           uint64
	CommerceID       uint64
	BranchID         uint64
	CampaignID       uint64
	AwardID          uint64
	CashBack         uint64
	RedeemedAwardID  uint64
	RedeemedCashBack float64
	IsPointsRedeem   bool
}
