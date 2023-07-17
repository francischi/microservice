package middleWare

import (
	// "fmt"
	"mainServer/pkg/base"
	"github.com/gin-gonic/gin"
	"mainServer/pkg/tokenModule"
)

func InitJwtMiddleWare()(middleWare *JwtMiddleWare){
	var JwtMiddleWare JwtMiddleWare

	tokenController := tokenModule.InitTokenController()
	JwtMiddleWare.tokenController = *tokenController
 	return &JwtMiddleWare
}

type JwtMiddleWare struct {
	base.MiddleWare
	tokenController tokenModule.TokenController
}

func (m *JwtMiddleWare) ConfirmToken(g *gin.Context){
	header := g.Request.Header["Bearer-Token"]
	if len(header)==0{
		m.InvaliAugument(g,"token_required")
		return
	}
	jwtToken := g.Request.Header["Bearer-Token"][0]

	val,err  := m.tokenController.IsValidJwt(jwtToken)
	if err!=nil{
		m.SystemError(g,err.Error())
		return
	}
	if !val {
		m.InvaliAugument(g,"invalid_token")
		return
	}

	val,err  = m.tokenController.IsJwtInTime(jwtToken)
	if err!=nil{
		m.SystemError(g,err.Error())
		return
	}
	if !val {
		m.InvaliAugument(g,"token_expired")
		return
	}
	
	g.Next()
}