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
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param   user   body   dto.CreateUserReq   true   "Información del usuario a crear"
// @Success 200 {object} domain.User "Usuario creado exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router / [post]
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
// @Tags pruchase
// @Accept  json
// @Produce  json
// @Param   user   body   dto.RegisterPurchaseReq   true   "Información de la compra a registrada"
// @Success 200 {object} domain.Purchase "Compra registrada exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /register-purchase [post]
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
