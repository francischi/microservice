//+build wireinject

package tokenModule

import (
	"github.com/google/wire"
)

func InitTokenController() *TokenController{
	wire.Build(
		NewTokenController , 
		NewTokenService , 
	)
	return &TokenController{}
}