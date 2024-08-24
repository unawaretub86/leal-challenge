package services

import "github.com/unawaretub86/leal-challenge/src/domain/ports"

type LealService struct {
	repository ports.RepositoryPort
}

func NewLealService(repository ports.RepositoryPort) *LealService {
	return &LealService{
		repository,
	}
}
