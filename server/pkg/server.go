package pkg

import (
	"context"
	"server/pkg/memberModule"
	"server/pkg/tokenModule"
	pb "server/proto"
)

type MemberService struct{
	pb.UnimplementedMemberServiceServer
}

func (s *MemberService)Testing(ctx context.Context , request *pb.TestRequest)( *pb.TestResponse , error){

	return &pb.TestResponse{
		Name: "this is name",
		Heigh: request.GetNum1() * request.GetNum2()* request.GetNum3(),
		Weight: request.GetNum1() + request.GetNum2() + request.GetNum3(),
	}, nil
}

func (s *MemberService)Create(ctx context.Context , request *pb.CreateReq)( *pb.CreateRes ,error){
	res , err := memberModule.InitMemberController().Create(ctx , request)
	return res,err
}

func (s *MemberService)LogIn(ctx context.Context , request *pb.LogInReq)( *pb.LogInRes , error){
	res ,err := memberModule.InitMemberController().LogIn(ctx , request)
	return res ,err
}

func (s *MemberService)ChangePwd(ctx context.Context , request *pb.ChangePwdReq)( *pb.ChangePwdRes , error){
	res , err := memberModule.InitMemberController().ChangePwd(ctx , request)
	return res,err
}

func (s *MemberService)IsValidJwt(ctx context.Context , request *pb.JwtToken)( *pb.IsValidJwtRes , error){
	res , err := tokenModule.InitTokenController().IsValidJwt(ctx ,request)
	return res,err
}

func (s *MemberService)IsJwtInTime(ctx context.Context , request *pb.JwtToken)( *pb.IsJwtInTimeRes , error){
	res , err := tokenModule.InitTokenController().IsJwtInTime(ctx ,request)
	return res,err
}