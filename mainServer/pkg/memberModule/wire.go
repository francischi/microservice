//+build wireinject

package memberModule

import (
	"github.com/google/wire"
)

func InitMemberController() *MemberController{
	wire.Build(
		NewMemberController , 
		NewMemberService , 
	)
	return &MemberController{}
}