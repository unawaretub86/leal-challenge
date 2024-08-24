package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type RepositoryPort interface {
	CreateUser(domain.User) (*domain.User, error)
	CreateCommerce(domain.Commerce) (*domain.Commerce, error)
	CreateBranch(domain.Branch) (*domain.Branch, error)
}
