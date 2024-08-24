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

func (s *LealService) CreateCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
	return s.repository.CreateCommerce(commerce)
}

func (s *LealService) CreateBranch(branch domain.Branch) (*domain.Branch, error) {
	return s.repository.CreateBranch(branch)
}

func (s *LealService) CreateCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return s.repository.CreateCampaign(campaign)
}

func (s *LealService) GetCommerceCampaigns(id uint64) (domain.Campaigns, error) {
	return s.repository.GetCommerceCampaigns(id)
}

func (s *LealService) GetBranchCampaigns(id uint64) (domain.Campaigns, error) {
	return s.repository.GetBranchCampaigns(id)
}
