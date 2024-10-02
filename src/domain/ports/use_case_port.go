package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type UseCasePort interface {
	GetUser(uint64) (*domain.User, error)
	GetCampaign(uint64) (*domain.Campaign, error)

	CalculateCashback(*domain.Purchase, domain.Campaign) error
	CalculatePoints(*domain.Purchase, domain.Campaign) error

	RedeemPoints(domain.Redeem) (*domain.Redeem, error)
	RedeemCashBack(domain.Redeem) (*domain.Redeem, error)
}
