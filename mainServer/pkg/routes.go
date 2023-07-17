package pkg

import (
	"github.com/gin-gonic/gin"
	"mainServer/pkg/memberModule"
)

func SetRouter(g *gin.Engine) {

	const baseGroup string = "/api"
	memberModule.SetRoute(g ,baseGroup)
}