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

func(m *MemberController) LogIn(g *gin.Context){
	var logInReq pb.LogInReq

	g.Bind(&logInReq)
	resp , err := m.MemberService.LogIn(&logInReq)
	if err!=nil{
		m.HandleError(g,err)
	}else{
		m.SuccessRes(g,resp.Token)
	}
}

func(m *MemberController) ChangePwd(g *gin.Context){
	var changePwdReq pb.ChangePwdReq

	g.Bind(&changePwdReq)
	_ , err := m.MemberService.ChangePwd(&changePwdReq)
	if err!=nil{
		m.HandleError(g,err)
	}else{
		m.SuccessRes(g,nil)
	}	
}