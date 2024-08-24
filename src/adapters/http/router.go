package http

import (
	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/leal-challenge/src/domain/ports"
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
	lealGroup := g.Group("/leal")

	userGroup := lealGroup.Group("/user")
	userGroup.POST("", r.CreateUser)
}
