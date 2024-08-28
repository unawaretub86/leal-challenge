package domain

type Award struct {
	ID          uint64 `gorm:"primaryKey"`
	CampaignID  uint64
	CommerceID  uint64
	Description string
	PointsCost  uint64
	Value       float64
}
