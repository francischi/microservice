package memberModule

import (
	// "fmt"
	pb "mainServer/proto"
	"mainServer/pkg/base"
	"github.com/gin-gonic/gin"
)

type MemberController struct{
	MemberService MemberService
	base.Controller
}

func NewMemberController(memberService *MemberService) *MemberController{
	return &MemberController{
		MemberService : *memberService,
	}
}

func(m *MemberController) Create(g *gin.Context){

	var createReq pb.CreateReq

	g.Bind(&createReq)
	_ , err := m.MemberService.Create(&createReq)
	if err!=nil{
		m.HandleError(g,err)
	}else{
		m.SuccessRes(g,nil)
	}
	
}

