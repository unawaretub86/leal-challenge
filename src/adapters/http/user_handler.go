package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/leal-challenge/src/adapters/http/dto"
	"github.com/unawaretub86/leal-challenge/utils/response"
)

func (r *LealRouter) CreateUser(c *gin.Context) {
	req := dto.CreateUserReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewUser(req)

	user, err := r.LealService.CreateUser(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixUser, err)
		return
	}

	response.EndWithStatus(c, http.StatusOK, user)
}
