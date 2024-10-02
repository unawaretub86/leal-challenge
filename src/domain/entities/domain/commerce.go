package domain

import "time"

type Commerce struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
