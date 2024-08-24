package dto

import (
	"time"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type CreateBranchDTO struct {
	Name       string    `json:"name" validate:"required"`
	CommerceID uint      `json:"commerce_id" validate:"required"`
	Address    string    `json:"address" validate:"required"`
	ActiveDate time.Time `json:"active_date" validate:"required"`
}

func NewBranch(dto CreateBranchDTO) *domain.Branch {
	return &domain.Branch{
		Name:       dto.Name,
		CommerceID: dto.CommerceID,
		Address:    dto.Address,
		StartDate:  dto.ActiveDate,
	}
}
