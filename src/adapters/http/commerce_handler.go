package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/leal-challenge/src/adapters/http/dto"
	"github.com/unawaretub86/leal-challenge/utils/response"
)

// Create Commerce
// @Summary Crear un nuevo comercio
// @Description Crea un comercio en el sistema utilizando la información proporcionada.
// @Tags Commerce
// @Accept  json
// @Produce  json
// @Param   commerce   body   dto.CreateCommerceDTO   true   "Información del comercio a crear"
// @Success 201 {object} domain.Commerce "comercio creado exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /commerce/ [post]
func (r *LealRouter) CreateCommerce(c *gin.Context) {
	req := dto.CreateCommerceDTO{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	params := dto.NewCommerce(req)

	commerce, err := r.LealService.CreateCommerce(*params)
	if err != nil {
		response.EndWithStatusError(c, http.StatusBadRequest, suffixCommerce, err)
		return
	}

	response.EndWithStatus(c, http.StatusCreated, commerce)
}

// Create Commerce branch
// @Summary Crear una nueva sucursal
// @Description Crea una sucursal en el sistema utilizando la información proporcionada.
// @Tags Branch
// @Accept  json
// @Produce  json
// @Param   branch   body   dto.CreateBranchDTO   true   "Información de la sucursal a crear"
// @Success 201 {object} domain.Branch "sucursal creada exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /commerce/branch [post]
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

// Create Campaign
// @Summary Crear una nueva campaña
// @Description Crea una campaña en el sistema utilizando la información proporcionada.
// @Tags Campaign
// @Accept  json
// @Produce  json
// @Param   campaign   body   dto.CreateCampaignDTO   true   "Información de la campaña a crear"
// @Success 201 {object} domain.Campaign "campaña creada exitosamente"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /commerce/campaign [post]
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

// Get Commerce Campaign
// @Summary Obtener campañas de un comercio
// @Description obtener campañas de un comercio en el sistema.
// @Tags Campaign
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Campaigns
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /commerce/campaign/:id [get]
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

// Get Branch Campaign
// @Summary Obtener campañas de una sucursal
// @Description obtener campañas de una sucursal en el sistema.
// @Tags Campaign
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Campaigns
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /commerce/branch/campaign/:id [get]
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
