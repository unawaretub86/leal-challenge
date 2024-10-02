package dto

import "github.com/unawaretub86/leal-challenge/src/domain/entities/domain"

type (
	CreateUserReq struct {
		Name  string `json:"name"  validate:"required"`
		Email string `json:"email"  validate:"required"`
	}

	RegisterPurchaseReq struct {
		UserID       uint64  `json:"user_id" validate:"required"`
		Amount       float64 `json:"amount" validate:"required"`
		CommerceID   uint64  `json:"commerce_id" validate:"required"`
		BranchID     uint64  `json:"branch_id" validate:"required"`
		CampaignID   uint64  `json:"campaign_id" validate:"required"`
		RedeemPoints bool    `json:"redeem_points" validate:"required"`
	}
)

func NewUser(userReq CreateUserReq) *domain.User {
	return &domain.User{
		Name:  userReq.Name,
		Email: userReq.Email,
	}
}

func NewPurchase(purchaseReq RegisterPurchaseReq) *domain.Purchase {
	return &domain.Purchase{
		UserID:       purchaseReq.UserID,
		Amount:       purchaseReq.Amount,
		CommerceID:   purchaseReq.CommerceID,
		BranchID:     purchaseReq.BranchID,
		CampaignID:   purchaseReq.CampaignID,
		RedeemPoints: purchaseReq.RedeemPoints,
	}
}
