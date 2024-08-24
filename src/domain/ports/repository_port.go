package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type RepositoryPort interface {
	CreateUser(domain.User) (*domain.User, error)
}
