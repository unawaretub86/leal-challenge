package domain

import "time"

type User struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email" gorm:"uniqueIndex"`
	Points    []Points   `gorm:"foreignKey:UserID"`
	Cashbacks []Cashback `gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
