package usecases

import (
	"errors"
	"time"

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

	purchase.EarnedCashBack = int(cashback)

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

	purchase.EarnedPoints = int(points)

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

func (useCase *LealUseCase) GetPurchase(id uint64) (*domain.Purchase, error) {
	return useCase.usecase.GetPurchase(id)
}

func (useCase *LealUseCase) RedeemPoints(redeem *domain.Redeem, purchase domain.Purchase, campaign domain.Campaign, user domain.User) error {
	if campaign.EndDate.Before(time.Now()) {
		return errors.New("campaña inactiva")
	}

	if purchase.CommerceID != campaign.CommerceID {
		return errors.New("no puede redimir puntos en un comercio diferente al que los adquirió")
	}

	var award *domain.Award
	for _, a := range campaign.Award {
		if a.ID == redeem.AwardID {
			award = &a
			break
		}
	}

	if user.Points >= int(award.PointsCost) {
		redeem.RedeemedAwardID = award.ID
		redeem.RedeemedPoints = award.PointsCost

		return nil
	}

	return errors.New("puntos insuficientes")
}

func (useCase *LealUseCase) RedeemCashBack(redeem *domain.Redeem, campaign domain.Campaign, user domain.User) error {
	if campaign.EndDate.Before(time.Now()) {
		return errors.New("campaña inactiva")
	}

	if user.Cashback >= int(redeem.CashBack) {
		redeem.RedeemedCashBack = float64(redeem.CashBack)

		return nil
	}

	return errors.New("no hay cashback")
}

func validateBonusFactor(bonusFactor float64) error {
	if bonusFactor <= 0 {
		return errors.New("bonus factor must be greater than 0")
	}

	return nil
}
