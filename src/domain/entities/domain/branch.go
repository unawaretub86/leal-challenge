package domain

import "time"

type Branch struct {
	ID         uint64 `gorm:"primaryKey"`
	Name       string
	CommerceID uint64
	Address    string
	StartDate  time.Time
	EndDate    *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
