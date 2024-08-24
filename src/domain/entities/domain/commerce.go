package domain

import "time"

type Commerce struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `json:"name" validate:"required"`
	Points    []Points   `gorm:"foreignKey:CommerceID"`
	Cashbacks []Cashback `gorm:"foreignKey:CommerceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}