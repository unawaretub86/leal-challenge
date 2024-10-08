package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type UseCasePort interface {
	// Get
	GetUser(uint64) (*domain.User, error)
	GetCampaign(uint64) (*domain.Campaign, error)
	GetPurchase(uint64) (*domain.Purchase, error)

	// Post
	CalculateCashback(*domain.Purchase, domain.Campaign) error
	CalculatePoints(*domain.Purchase, domain.Campaign) error
	RedeemPoints(*domain.Redeem, domain.Purchase, domain.Campaign, domain.User) error
	RedeemCashBack(*domain.Redeem, domain.Campaign, domain.User) error
}
