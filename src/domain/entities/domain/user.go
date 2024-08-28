package domain

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Points    int
	Cashback  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
