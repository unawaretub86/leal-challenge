package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
)

type LealRepository struct {
	db *gorm.DB
}

func NewLealRepository(db *gorm.DB) *LealRepository {
	return &LealRepository{
		db: db,
	}
}

func (r *LealRepository) CreateCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
	if err := r.db.Create(&commerce).Error; err != nil {
		return nil, err
	}

	return &commerce, nil
}

func (r *LealRepository) CreateBranch(branch domain.Branch) (*domain.Branch, error) {
	if err := r.db.Create(&branch).Error; err != nil {
		return nil, err
	}

	return &branch, nil
}

func (r *LealRepository) CreateCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	forAll, branchCampaign, err := r.isActiveCampaign(*campaign.BranchID)
	if err != nil {
		return nil, err
	}

	if forAll || branchCampaign {
		return nil, errors.New("campaign is already active")
	}

	if err := r.db.Create(&campaign).Error; err != nil {
		return nil, err
	}

	return &campaign, nil
}

func (r *LealRepository) GetCommerceCampaigns(id uint64) (domain.Campaigns, error) {
	var resultData domain.Campaigns

	result := r.db.Model(resultData).
		Where("commerce_id = ?", id).
		Find(&resultData)

	if result.Error != nil {
		return nil, result.Error
	}

	return resultData, nil
}

func (r *LealRepository) GetBranchCampaigns(id uint64) (domain.Campaigns, error) {
	var resultData domain.Campaigns

	result := r.db.Model(resultData).
		Where("branch_id = ?", id).
		Find(&resultData)

	if result.Error != nil {
		return nil, result.Error
	}

	return resultData, nil
}

func (r *LealRepository) GetCampaign(id uint64) (*domain.Campaign, error) {
	resultData, err := r.getCampaign(id)
	if err.Error != nil {
		return nil, err.Error
	}

	return resultData, nil
}

func (r *LealRepository) getCampaign(id uint64) (*domain.Campaign, *gorm.DB) {
	campaign := &domain.Campaign{}

	result := r.db.
		Where("id = ?", id).
		Take(&campaign)

	return campaign, result
}

func (r *LealRepository) isActiveCampaign(branchID uint64) (bool, bool, error) {
	campaign := domain.Campaign{}

	result := r.db.Model(campaign).
		Where("branch_id = ?", branchID).
		Order("created_at DESC").
		First(&campaign)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, false, result.Error
	}

	if time.Now().After(campaign.EndDate) {
		return false, false, nil
	}

	if campaign.IsForAll {
		return true, false, nil
	}

	return false, true, nil
}
