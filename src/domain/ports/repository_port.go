package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type RepositoryPort interface {
	// commerce methods
	// Post
	CreateUser(domain.User) (*domain.User, error)
	CreateCommerce(domain.Commerce) (*domain.Commerce, error)
	CreateBranch(domain.Branch) (*domain.Branch, error)
	CreateCampaign(domain.Campaign) (*domain.Campaign, error)

	// Get methods
	GetCommerceCampaigns(id uint64) (domain.Campaigns, error)
	GetBranchCampaigns(id uint64) (domain.Campaigns, error)

	// User methods
	//  Post
	RegisterPurchase(purchase domain.Purchase) (*domain.Purchase, error)
}
