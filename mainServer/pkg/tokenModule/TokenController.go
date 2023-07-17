package tokenModule

import (
	// "fmt"
	pb "mainServer/proto"
	"mainServer/pkg/base"
)

type TokenController struct{
	TokenService TokenService
	base.Controller
}

func NewTokenController(tokenService *TokenService) *TokenController{
	return &TokenController{
		TokenService : *tokenService,
	}
}

func(t *TokenController) IsJwtInTime(token string)(bool , error){
	var jwtToken pb.JwtToken
	jwtToken.Token = token

	resp , err := t.TokenService.IsJwtInTime(&jwtToken)
	if err!=nil{
		return false,err
	}else{
		return resp.Success,nil
	}
}

func(t *TokenController) IsValidJwt(token string)(bool , error){
	var jwtToken pb.JwtToken
	jwtToken.Token = token

	resp , err := t.TokenService.IsValidJwt(&jwtToken)
	if err!=nil{
		return false,err
	}else{
		return resp.Success,nil
	}
}