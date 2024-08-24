package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/leal-challenge/src/adapters/http/dto"
	"github.com/unawaretub86/leal-challenge/utils/response"
)

func (r *LealRouter) CreateCommerce(c *gin.Context) {
	req := dto.CreateCommerceDTO{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewCommerce(req)

	comerce, err := r.LealService.CreateCommerce(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, comerce)
}

func (r *LealRouter) CreateBranch(c *gin.Context) {
	req := dto.CreateBranchDTO{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewBranch(req)

	branch, err := r.LealService.CreateBranch(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, branch)
}

func (r *LealRouter) CreateCampaign(c *gin.Context) {
	req := dto.CreateCampaignDTO{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params, err := dto.NewCampaign(req)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	campaign, err := r.LealService.CreateCampaign(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, campaign)
}

func (r *LealRouter) GetCommerceCampaigns(c *gin.Context) {
	id := c.Param("id")
	commerceID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	campaigns, err := r.LealService.GetCommerceCampaigns(commerceID)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusOK, campaigns)
}

func (r *LealRouter) GetBranchCampaigns(c *gin.Context) {
	id := c.Param("id")
	branchID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	campaigns, err := r.LealService.GetBranchCampaigns(branchID)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusOK, campaigns)
}
