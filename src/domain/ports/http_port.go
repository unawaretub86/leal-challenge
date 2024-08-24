package ports

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type LealPort interface {
	CreateUser(domain.User) (*domain.User, error)
}
