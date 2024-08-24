package domain

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email" gorm:"uniqueIndex"`
	Points    []Points   `gorm:"foreignKey:UserID"`
	Cashbacks []Cashback `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
