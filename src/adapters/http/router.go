package http

import (
	"github.com/gin-gonic/gin"
	
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
)

type LealRouter struct {
	LeaService ports.LealPort
}

func NewRouter(LeaServicePorts ports.LealPort) *LealRouter {
	return &LealRouter{
		LeaService: LeaServicePorts,
	}
}

func (r *LealRouter) SetRoutes(g *gin.Engine) {
	_ = g.Group("/")
}
