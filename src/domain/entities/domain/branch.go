package domain

import "time"

type Branch struct {
	ID         uint64     `gorm:"primaryKey"`
	Name       string     `json:"name" validate:"required"`
	CommerceID uint64     `json:"commerce_id" validate:"required"`
	Address    string     `json:"address"`
	StartDate  time.Time  `json:"start_date"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
