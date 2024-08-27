package usecases

import (
	"errors"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

type LealUseCase struct {
	usecase ports.LealPort
}

func NewLealUseCase(usecase ports.RepositoryPort) *LealUseCase {
	return &LealUseCase{
		usecase,
	}
}

func (useCase LealUseCase) CalculateCashback(purchase *domain.Purchase, campaign domain.Campaign) error {
	if campaign.BonusFactor <= 0 {
		return errors.New("bonus factor is 0")
	}

	purchase.EarnedCashBack = int(purchase.Amount / campaign.BonusFactor)

	return nil
}

func (useCase LealUseCase) CalculatePoints(purchase *domain.Purchase, campaign domain.Campaign) error {
	if campaign.BonusFactor <= 0 {
		return errors.New("bonus factor is 0")
	}

	purchase.EarnedPoints = int(purchase.Amount / campaign.BonusFactor)

	return nil
}

func (useCase *LealUseCase) GetCampaign(id uint64) (*domain.Campaign, error) {
	return useCase.usecase.GetCampaign(id)
}

func (useCase *LealUseCase) GetUser(id uint64) (*domain.User, error) {
	return useCase.usecase.GetUser(id)
}
