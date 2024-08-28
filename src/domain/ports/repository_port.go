package ports

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type RepositoryPort interface {
	// commerce methods
	// Post
	CreateUser(domain.User) (*domain.User, error)
	CreateCommerce(domain.Commerce) (*domain.Commerce, error)
	CreateBranch(domain.Branch) (*domain.Branch, error)
	CreateCampaign(domain.Campaign) (*domain.Campaign, error)

	// Get
	GetCommerceCampaigns(uint64) (domain.Campaigns, error)
	GetBranchCampaigns(uint64) (domain.Campaigns, error)
	GetCampaign(uint64) (*domain.Campaign, error)

	// User methods
	//  Post
	RegisterPurchase(domain.Purchase) (*domain.Purchase, error)
	UpdateUser(uint64, domain.User) (*domain.User, error)

	// Get
	GetUser(uint64) (*domain.User, error)
}
