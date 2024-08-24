package http

import (
	"net/http"

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

	response.EndWithStatus(c, http.StatusOK, comerce)
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

	response.EndWithStatus(c, http.StatusOK, branch)
}
