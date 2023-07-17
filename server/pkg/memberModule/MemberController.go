package memberModule

import (
	// "fmt"
	"context"
   "server/pkg/memberModule/dtos" 
   "google.golang.org/grpc/status"
   "server/pkg/base"
   pb "server/proto"
)

type MemberController struct{
	MemberService MemberService
	// pb.UnimplementedMemberServiceServer
	base.Controller
}

func NewMemberController(memberService *MemberService) *MemberController{
	return &MemberController{
		MemberService : *memberService,
	}
}

func (c *MemberController)Testing( ctx context.Context , request *pb.TestRequest)( *pb.TestResponse , error){

	return &pb.TestResponse{
		Name: "this is name",
		Heigh: request.GetNum1() * request.GetNum2()* request.GetNum3(),
		Weight: request.GetNum1() + request.GetNum2() + request.GetNum3(),
	}, nil
}

func (c *MemberController)Create(ctx context.Context , req *pb.CreateReq) (*pb.CreateRes , error) {

	var member dtos.CreateMemberDto
	member.Name = req.Name
	member.Email = req.Email
	member.Gender = req.Gender
	member.Password = req.Password

	err := c.MemberService.Create(&member)
	if err!=nil{
		return &pb.CreateRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	return &pb.CreateRes{Success: true , ErrMsg: ""} ,nil
}

func (c *MemberController) LogIn(ctx context.Context , req *pb.LogInReq)(*pb.LogInRes ,error){

	var member dtos.LogInDto
	member.Account = req.Account
	member.Password = req.Password
	
	token , err := c.MemberService.LogIn(&member)
	// return token , err
	if err!=nil{
		return &pb.LogInRes{
			Success: false ,
			Token: "" ,  
			ErrMsg: err.Error(),
		},
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	return &pb.LogInRes{Success: true ,Token: token, ErrMsg: ""} ,nil
}

func (c *MemberController ) ChangePwd(ctx context.Context , req *pb.ChangePwdReq)(*pb.ChangePwdRes ,error){
	
	var member dtos.ChangePwdDto
	member.MemberId = req.MemberId
	member.ConfirmPassword = req.ConfirmPassword
	member.OldPassword = req.OldPassword
	member.NewPassword = req.NewPassword

	err := c.MemberService.ChangePwd(&member)
	if err!=nil{
		return &pb.ChangePwdRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	return &pb.ChangePwdRes{Success: true , ErrMsg: ""} ,nil
}