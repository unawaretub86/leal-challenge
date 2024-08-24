package dto

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type (
	CreateUserReq struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func NewUser(userReq CreateUserReq) *domain.User {
	return &domain.User{
		Name:  userReq.Name,
		Email: userReq.Email,
	}
}
