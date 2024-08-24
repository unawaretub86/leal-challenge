package services

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

type LealService struct {
	repository ports.RepositoryPort
}

func NewLealService(repository ports.RepositoryPort) *LealService {
	return &LealService{
		repository,
	}
}

func (s *LealService) CreateUser(user domain.User) (*domain.User, error) {
	return s.repository.CreateUser(user)
}
