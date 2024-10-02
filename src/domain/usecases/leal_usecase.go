package usecases

import (
	"errors"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

const (
	PERCENT = "percentage"
	DOUBLE  = "double"
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
	if err := validateBonusFactor(campaign.BonusFactor); err != nil {
		return err
	}

	cashback, err := useCase.calculateReward(purchase.Amount, campaign.BonusFactor, campaign.BonusType, campaign.MinPurchase)
	if err != nil {
		return err
	}

	earnedCashBack := int(cashback)

	purchase.EarnedCashBack = earnedCashBack

	return nil
}

func (useCase LealUseCase) CalculatePoints(purchase *domain.Purchase, campaign domain.Campaign) error {
	if err := validateBonusFactor(campaign.BonusFactor); err != nil {
		return err
	}

	points, err := useCase.calculateReward(purchase.Amount, campaign.BonusFactor, campaign.BonusType, campaign.MinPurchase)
	if err != nil {
		return err
	}

	earnedPoints := int(points)

	purchase.EarnedPoints = earnedPoints

	return nil
}

func (useCase LealUseCase) calculateReward(amount float64, bonusFactor float64, bonusType string, minPurchase *float64) (float64, error) {
	earnedReward := amount / 1000

	switch bonusType {
	case PERCENT:
		if amount >= *minPurchase {
			return earnedReward + (earnedReward * (bonusFactor / 100)), nil
		}
	case DOUBLE:
		return earnedReward * bonusFactor, nil
	default:
		return 0, errors.New("invalid bonus type: " + bonusType)
	}

	return 0, nil
}

func (useCase *LealUseCase) GetCampaign(id uint64) (*domain.Campaign, error) {
	return useCase.usecase.GetCampaign(id)
}

func (useCase *LealUseCase) GetUser(id uint64) (*domain.User, error) {
	return useCase.usecase.GetUser(id)
}

func validateBonusFactor(bonusFactor float64) error {
	if bonusFactor <= 0 {
		return errors.New("bonus factor must be greater than 0")
	}
	return nil
}
