package ports

import (
	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type LealPort interface {
	// commerce methods
	// Post
	CreateUser(domain.User) (*domain.User, error)
	CreateCommerce(domain.Commerce) (*domain.Commerce, error)
	CreateBranch(domain.Branch) (*domain.Branch, error)
	CreateCampaign(domain.Campaign) (*domain.Campaign, error)

	// Get methods
	GetCommerceCampaigns(id uint64) (domain.Campaigns, error)
	GetBranchCampaigns(id uint64) (domain.Campaigns, error)
	GetCampaign(id uint64) (*domain.Campaign, error)
	GetPurchase(uint64) (*domain.Purchase, error)

	// User methods
	//  Post
	RegisterPurchase(domain.Purchase) (*domain.Purchase, error)
	Redeem(domain.Redeem) (*domain.Redeem, error)

	// Get
	GetUser(uint64) (*domain.User, error)
}
