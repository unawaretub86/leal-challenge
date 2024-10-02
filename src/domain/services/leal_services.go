package services

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

type LealService struct {
	repository ports.RepositoryPort
	usecases   ports.UseCasePort
}

func NewLealService(repository ports.RepositoryPort,
	usecases ports.UseCasePort) *LealService {
	return &LealService{
		repository,
		usecases,
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

func (s *LealService) GetCampaign(id uint64) (*domain.Campaign, error) {
	return s.repository.GetCampaign(id)
}

func (s *LealService) GetUser(id uint64) (*domain.User, error) {
	return s.repository.GetUser(id)
}

func (s *LealService) GetPurchase(id uint64) (*domain.Purchase, error) {
	return s.repository.GetPurchase(id)
}

func (s *LealService) RegisterPurchase(purchase domain.Purchase) (*domain.Purchase, error) {
	campaign, err := s.usecases.GetCampaign(purchase.CampaignID)
	if err != nil {
		return nil, err
	}

	if purchase.RedeemPoints {
		err = s.usecases.CalculatePoints(&purchase, *campaign)
		if err != nil {
			return nil, err
		}

		return s.repository.RegisterPurchase(purchase)
	}

	err = s.usecases.CalculateCashback(&purchase, *campaign)
	if err != nil {
		return nil, err
	}

	return s.repository.RegisterPurchase(purchase)
}

func (s *LealService) Redeem(redeem domain.Redeem) (*domain.Redeem, error) {
	campaign, err := s.usecases.GetCampaign(redeem.CampaignID)
	if err != nil {
		return nil, err
	}

	purchase, err := s.usecases.GetPurchase(redeem.PurchaseID)
	if err != nil {
		return nil, err
	}

	user, err := s.GetUser(redeem.UserID)
	if err != nil {
		return nil, err
	}

	if redeem.IsPointsRedeem {
		err = s.usecases.RedeemPoints(&redeem, *purchase, *campaign, *user)
		if err != nil {
			return nil, err
		}

		return s.repository.Redeem(redeem)
	}

	err = s.usecases.RedeemCashBack(&redeem, *campaign, *user)
	if err != nil {
		return nil, err
	}

	return s.repository.Redeem(redeem)
}
