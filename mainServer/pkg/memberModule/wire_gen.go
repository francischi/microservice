// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package memberModule

// Injectors from wire.go:

func InitMemberController() *MemberController {
	memberService := NewMemberService()
	memberController := NewMemberController(memberService)
	return memberController
}