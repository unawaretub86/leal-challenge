package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/leal-challenge/src/adapters/http/dto"
	"github.com/unawaretub86/leal-challenge/utils/response"
)

// CreateUser
// @Summary Crear un nuevo usuario
// @Description Crea un usuario en el sistema utilizando la información proporcionada.
// @Tags User
// @Accept  json
// @Produce  json
// @Param   user   body   dto.CreateUserReq   true   "Información del usuario a crear"
// @Success 201 {object} domain.User "Usuario creado exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /user/ [post]
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

// RegisterPurchase
// @Summary Registrar una nueva compra
// @Description Registra una nueva compra en el sistema utilizando la información proporcionada.
// @Tags Purchase
// @Accept  json
// @Produce  json
// @Param   purchase   body   dto.RegisterPurchaseReq   true   "Información de la compra a registrada"
// @Success 201 {object} domain.Purchase "Compra registrada exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /user/register-purchase [post]
func (r *LealRouter) RegisterPurchase(c *gin.Context) {
	req := dto.RegisterPurchaseReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewPurchase(req)

	purchase, err := r.LealService.RegisterPurchase(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixUser, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, purchase)
}

// Redeem points
// @Summary Registrar una redencion de puntos o cashback
// @Description Registrar una redencion de puntos o cashback.
// @Tags Redeem
// @Accept  json
// @Produce  json
// @Param   redeem   body   dto.RedeemReq   true   "Información para redimir"
// @Success 201 {object} domain.Redeem "Compra registrada exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /user/redeem [post]
func (r *LealRouter) Redeem(c *gin.Context) {
	req := dto.RedeemReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewRedeem(req)

	purchase, err := r.LealService.Redeem(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixUser, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, purchase)
}
