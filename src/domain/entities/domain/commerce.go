package domain

import "time"

type Commerce struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
