package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/leal-challenge/src/domain/entities/domain"
	"github.com/unawaretub86/leal-challenge/src/domain/services"
	mocks_test "github.com/unawaretub86/leal-challenge/test/mocks"
)

type mocks struct {
	repoDB    *mocks_test.MockRepositoryPort
	lealCases *mocks_test.MockUseCasePort
}

func setupMocks(t *testing.T) (mocks, *services.LealService) {
	ctrl := gomock.NewController(t)

	m := mocks{
		repoDB:    mocks_test.NewMockRepositoryPort(ctrl),
		lealCases: mocks_test.NewMockUseCasePort(ctrl),
	}

	lealService := services.NewLealService(m.repoDB, m.lealCases)
	return m, lealService
}

func TestCreateCampaignOK(t *testing.T) {
	m, lealService := setupMocks(t)

	campaign := domain.Campaign{
		Name: "test 1",
	}

	m.repoDB.EXPECT().CreateCampaign(campaign).Return(&campaign, nil)

	campaignCreated, err := lealService.CreateCampaign(campaign)
	assert.Nil(t, err)
	assert.NotNil(t, campaignCreated)
}

func TestCreateCampaignFail(t *testing.T) {
	m, lealService := setupMocks(t)

	campaign := domain.Campaign{}

	m.repoDB.EXPECT().CreateCampaign(campaign).Return(nil, errors.New("empty campaign"))

	campaignCreated, err := lealService.CreateCampaign(campaign)

	assert.NotNil(t, err)
	assert.Nil(t, campaignCreated)
}

func TestCreateBranchOK(t *testing.T) {
	m, lealService := setupMocks(t)

	branch := domain.Branch{
		Name:       "test 1",
		CommerceID: 1,
		Address:    "av siempre viva",
	}

	m.repoDB.EXPECT().CreateBranch(branch).Return(&branch, nil)

	branchCreated, err := lealService.CreateBranch(branch)

	assert.Nil(t, err)
	assert.NotNil(t, branchCreated)
}

func TestCreateBranchFail(t *testing.T) {
	m, lealService := setupMocks(t)

	branch := domain.Branch{}

	m.repoDB.EXPECT().CreateBranch(branch).Return(nil, errors.New("empty branch"))

	branchCreated, err := lealService.CreateBranch(branch)

	assert.NotNil(t, err)
	assert.Nil(t, branchCreated)
}

func TestCreateCommerceOK(t *testing.T) {
	m, lealService := setupMocks(t)

	commerce := domain.Commerce{
		Name: "test 1",
	}

	m.repoDB.EXPECT().CreateCommerce(commerce).Return(&commerce, nil)

	commerceCreated, err := lealService.CreateCommerce(commerce)

	assert.Nil(t, err)
	assert.NotNil(t, commerceCreated)
}

func TestCreateCommerceFail(t *testing.T) {
	m, lealService := setupMocks(t)

	commerce := domain.Commerce{}

	m.repoDB.EXPECT().CreateCommerce(commerce).Return(nil, errors.New("empty commerce"))

	commerceCreated, err := lealService.CreateCommerce(commerce)

	assert.NotNil(t, err)
	assert.Nil(t, commerceCreated)
}

func TestCreateUserOK(t *testing.T) {
	m, lealService := setupMocks(t)

	user := domain.User{
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.repoDB.EXPECT().CreateUser(user).Return(&user, nil)

	userCreated, err := lealService.CreateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, userCreated)
}

func TestCreateUserFail(t *testing.T) {
	m, lealService := setupMocks(t)

	user := domain.User{}

	m.repoDB.EXPECT().CreateUser(user).Return(nil, errors.New("empty user"))

	userCreated, err := lealService.CreateUser(user)

	assert.NotNil(t, err)
	assert.Nil(t, userCreated)
}

func TestGetCampaignOK(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	campaign := domain.Campaign{
		ID:   1,
		Name: "test 1",
	}

	m.repoDB.EXPECT().GetCampaign(id).Return(&campaign, nil)

	campaigns, err := lealService.GetCampaign(id)

	assert.Nil(t, err)
	assert.NotNil(t, campaigns)
}

func TestGetCampaignFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	m.repoDB.EXPECT().GetCampaign(id).Return(nil, errors.New("empty campaigns"))

	campaigns, err := lealService.GetCampaign(id)

	assert.NotNil(t, err)
	assert.Nil(t, campaigns)
}

func TestGetBranchCampaignsOK(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	campaign := domain.Campaign{
		ID:   1,
		Name: "test 1",
	}

	campaigns := domain.Campaigns{
		campaign,
	}

	m.repoDB.EXPECT().GetBranchCampaigns(id).Return(campaigns, nil)

	branchCampaigns, err := lealService.GetBranchCampaigns(id)

	assert.Nil(t, err)
	assert.NotNil(t, branchCampaigns)
}

func TestGetBranchCampaignsFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	m.repoDB.EXPECT().GetBranchCampaigns(id).Return(nil, errors.New("empty campaigns"))

	branchCampaigns, err := lealService.GetBranchCampaigns(id)

	assert.NotNil(t, err)
	assert.Nil(t, branchCampaigns)
}

func TestGetPurchaseOK(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	purchased := domain.Purchase{
		ID:           1,
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	m.repoDB.EXPECT().GetPurchase(id).Return(&purchased, nil)

	purchase, err := lealService.GetPurchase(id)

	assert.Nil(t, err)
	assert.NotNil(t, purchase)
}

func TestGetPurchaseFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	m.repoDB.EXPECT().GetPurchase(id).Return(nil, errors.New("empty purchase"))

	purchase, err := lealService.GetPurchase(id)

	assert.NotNil(t, err)
	assert.Nil(t, purchase)
}

func TestGetCommerceCampaignsOK(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	campaign := domain.Campaign{
		ID:   1,
		Name: "test 1",
	}

	campaigns := domain.Campaigns{
		campaign,
	}

	m.repoDB.EXPECT().GetCommerceCampaigns(id).Return(campaigns, nil)

	commerceCampaigns, err := lealService.GetCommerceCampaigns(id)

	assert.Nil(t, err)
	assert.NotNil(t, commerceCampaigns)
}

func TestGetCommerceCampaignsFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	m.repoDB.EXPECT().GetCommerceCampaigns(id).Return(nil, errors.New("empty commerce campaigns"))

	purchase, err := lealService.GetCommerceCampaigns(id)

	assert.NotNil(t, err)
	assert.Nil(t, purchase)
}

func TestGetUserOK(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1
	user := domain.User{
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.repoDB.EXPECT().GetUser(id).Return(&user, nil)

	userResult, err := lealService.GetUser(id)

	assert.Nil(t, err)
	assert.NotNil(t, userResult)
}

func TestGetUserFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1

	m.repoDB.EXPECT().GetUser(id).Return(nil, errors.New("empty user"))

	userResult, err := lealService.GetUser(id)

	assert.NotNil(t, err)
	assert.Nil(t, userResult)
}

func TestCreatePurchaseOK(t *testing.T) {
	m, lealService := setupMocks(t)

	purchase := &domain.Purchase{
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	purchased := domain.Purchase{
		ID:           0,
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	campaign := domain.Campaign{
		Name: "test 1",
	}

	m.lealCases.EXPECT().GetCampaign(purchase.CampaignID).Return(&campaign, nil)
	m.lealCases.EXPECT().CalculatePoints(purchase, campaign).Return(nil)
	m.repoDB.EXPECT().RegisterPurchase(purchased).Return(&purchased, nil)

	purchaseCreated, err := lealService.RegisterPurchase(*purchase)

	assert.Nil(t, err)
	assert.NotNil(t, purchaseCreated)
}

func TestCreatePurchaseFail(t *testing.T) {
	m, lealService := setupMocks(t)

	purchase := &domain.Purchase{
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	purchased := domain.Purchase{
		ID:           0,
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	campaign := domain.Campaign{
		Name: "test 1",
	}

	m.lealCases.EXPECT().GetCampaign(purchase.CampaignID).Return(&campaign, nil)
	m.lealCases.EXPECT().CalculatePoints(purchase, campaign).Return(nil)
	m.repoDB.EXPECT().RegisterPurchase(purchased).Return(nil, errors.New("error creating purchase"))

	purchaseCreated, err := lealService.RegisterPurchase(*purchase)

	assert.NotNil(t, err)
	assert.Nil(t, purchaseCreated)
}

func TestCreateRedeemPointOK(t *testing.T) {
	m, lealService := setupMocks(t)
	var id uint64 = 1

	redeem := domain.Redeem{
		CommerceID:     1,
		UserID:         1,
		BranchID:       1,
		CampaignID:     1,
		AwardID:        1,
		PurchaseID:     1,
		IsPointsRedeem: true,
	}

	purchase := domain.Purchase{
		ID:           1,
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	award := domain.Award{
		ID:          1,
		CampaignID:  1,
		CommerceID:  1,
		Description: "test",
		PointsCost:  10,
		Value:       1,
	}

	campaign := domain.Campaign{
		ID:          1,
		Name:        "test 1",
		CommerceID:  1,
		BranchID:    &id,
		BonusFactor: 2,
		Award: []domain.Award{
			award,
		},
	}

	user := domain.User{
		ID:    1,
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.lealCases.EXPECT().GetCampaign(redeem.CampaignID).Return(&campaign, nil)
	m.lealCases.EXPECT().GetPurchase(redeem.CampaignID).Return(&purchase, nil)
	m.repoDB.EXPECT().GetUser(id).Return(&user, nil)

	m.lealCases.EXPECT().RedeemPoints(&redeem, purchase, campaign, user).Return(nil)
	m.repoDB.EXPECT().Redeem(redeem).Return(&redeem, nil)

	purchaseCreated, err := lealService.Redeem(redeem)

	assert.Nil(t, err)
	assert.NotNil(t, purchaseCreated)
}

func TestCreateRedeemPointFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1

	redeem := domain.Redeem{
		CommerceID:     1,
		UserID:         1,
		BranchID:       1,
		CampaignID:     1,
		AwardID:        1,
		PurchaseID:     1,
		IsPointsRedeem: true,
	}

	purchase := domain.Purchase{
		ID:           1,
		Amount:       100,
		CommerceID:   1,
		UserID:       1,
		BranchID:     1,
		CampaignID:   1,
		RedeemPoints: true,
	}

	award := domain.Award{
		ID:          1,
		CampaignID:  1,
		CommerceID:  1,
		Description: "test",
		PointsCost:  10,
		Value:       1,
	}

	campaign := domain.Campaign{
		ID:          1,
		Name:        "test 1",
		CommerceID:  1,
		BranchID:    &id,
		BonusFactor: 2,
		Award: []domain.Award{
			award,
		},
	}

	user := domain.User{
		ID:    1,
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.lealCases.EXPECT().GetCampaign(redeem.CampaignID).Return(&campaign, nil)
	m.lealCases.EXPECT().GetPurchase(redeem.CampaignID).Return(&purchase, nil)
	m.repoDB.EXPECT().GetUser(id).Return(&user, nil)

	m.lealCases.EXPECT().RedeemPoints(&redeem, purchase, campaign, user).Return(nil)

	m.repoDB.EXPECT().Redeem(redeem).Return(nil, errors.New("error redeem"))

	purchaseCreated, err := lealService.Redeem(redeem)
	assert.NotNil(t, err)
	assert.Nil(t, purchaseCreated)
}

func TestCreateRedeemCashBackOK(t *testing.T) {
	m, lealService := setupMocks(t)
	var id uint64 = 1

	redeem := domain.Redeem{
		CommerceID:     1,
		UserID:         1,
		BranchID:       1,
		CampaignID:     1,
		AwardID:        1,
		PurchaseID:     1,
		IsPointsRedeem: false,
	}

	campaign := domain.Campaign{
		ID:          1,
		Name:        "test 1",
		CommerceID:  1,
		BranchID:    &id,
		BonusFactor: 2,
	}

	user := domain.User{
		ID:    1,
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.lealCases.EXPECT().GetCampaign(redeem.CampaignID).Return(&campaign, nil)
	m.repoDB.EXPECT().GetUser(id).Return(&user, nil)

	m.lealCases.EXPECT().RedeemCashBack(&redeem, campaign, user).Return(nil)
	m.repoDB.EXPECT().Redeem(redeem).Return(&redeem, nil)

	purchaseCreated, err := lealService.Redeem(redeem)

	assert.Nil(t, err)
	assert.NotNil(t, purchaseCreated)
}

func TestCreateRedeemCashBackFail(t *testing.T) {
	m, lealService := setupMocks(t)

	var id uint64 = 1

	redeem := domain.Redeem{
		CommerceID:     1,
		UserID:         1,
		BranchID:       1,
		CampaignID:     1,
		AwardID:        1,
		PurchaseID:     1,
		IsPointsRedeem: false,
	}

	campaign := domain.Campaign{
		ID:          1,
		Name:        "test 1",
		CommerceID:  1,
		BranchID:    &id,
		BonusFactor: 2,
	}

	user := domain.User{
		ID:    1,
		Name:  "jhon doe",
		Email: "john.doe1@example.com",
	}

	m.lealCases.EXPECT().GetCampaign(redeem.CampaignID).Return(&campaign, nil)
	m.repoDB.EXPECT().GetUser(id).Return(&user, nil)

	m.lealCases.EXPECT().RedeemCashBack(&redeem, campaign, user).Return(nil)

	m.repoDB.EXPECT().Redeem(redeem).Return(nil, errors.New("error redeem"))

	purchaseCreated, err := lealService.Redeem(redeem)

	assert.NotNil(t, err)
	assert.Nil(t, purchaseCreated)
}
