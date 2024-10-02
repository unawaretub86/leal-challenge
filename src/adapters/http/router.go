package http

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/unawaretub86/leal-challenge/cmd/docs"

	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

const (
	suffixErr      = "Error"
	suffixUser     = "User"
	suffixCommerce = "Commerce"
)

type LealRouter struct {
	LealService ports.LealPort
}

func NewRouter(LealServicePorts ports.LealPort) *LealRouter {
	return &LealRouter{
		LealService: LealServicePorts,
	}
}

func (r *LealRouter) SetRoutes(g *gin.Engine) {
	r.SetTestRoutes(g)

	lealGroup := g.Group("/leal")

	userGroup := lealGroup.Group("/user")
	userGroup.POST("", r.CreateUser)
	userGroup.POST("/register-purchase", r.RegisterPurchase)

	commerceGroup := lealGroup.Group("/commerce")
	commerceGroup.POST("", r.CreateCommerce)
	commerceGroup.POST("/branch", r.CreateBranch)
	commerceGroup.POST("/campaign", r.CreateCampaign)
	commerceGroup.GET("/campaign/:id", r.GetCommerceCampaigns)
	commerceGroup.GET("/branch/campaign/:id", r.GetBranchCampaigns)
}

// @title Leal challenge
// @version 1.0
// @description Esta api es para poder solucionar el reto tecnico de leal
// @termsOfService http://swagger.io/terms/

// @contact.name Esteban Gomez
// @contact.email unawaretub86@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
func (r *LealRouter) SetTestRoutes(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
