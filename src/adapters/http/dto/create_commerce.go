package dto

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type CreateCommerceDTO struct {
	Name string `json:"name" validate:"required"`
}

func NewCommerce(commerce CreateCommerceDTO) *domain.Commerce {
	return &domain.Commerce{
		Name: commerce.Name,
	}
}
