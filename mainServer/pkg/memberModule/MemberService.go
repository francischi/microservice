package memberModule

import (
	"fmt"
	"time"
	"context"
	pb "mainServer/proto"
	"mainServer/pkg/base"
)

type MemberService struct{
	base.Service
	ServicName  string 
}

func NewMemberService () *MemberService{
	var memberService MemberService
	memberService.ServicName = "memberService"
	return &memberService
} 

func (m *MemberService) Create(req *pb.CreateReq)(*pb.CreateRes , error){
	var resp  *pb.CreateRes

	serviceAddress ,err := m.FindService(m.ServicName)
	if err!=nil{
		return resp , err
	}

	conn , err := m.CreateConn(serviceAddress)
	if err!=nil{
		return resp , err
	}
	defer conn.Close()

	memberServiceClient := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err = memberServiceClient.Create(ctx , req)
	if err!=nil{
		return resp , err
	}else {
		return resp , nil
	}
}

func (m *MemberService) LogIn(req *pb.LogInReq)(*pb.LogInRes , error){
	var resp  *pb.LogInRes

	serviceAddress ,err := m.FindService(m.ServicName)
	fmt.Println(serviceAddress)
	if err!=nil{
		return resp , err
	}
	conn , err := m.CreateConn(serviceAddress)
	if err!=nil{
		return resp , err
	}
	defer conn.Close()

	memberServiceClient := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err = memberServiceClient.LogIn(ctx , req)
	if err!=nil{
		return resp , err
	}else {
		return resp , nil
	}
}

func (m *MemberService) ChangePwd(req *pb.ChangePwdReq)(*pb.ChangePwdRes , error){
	var resp  *pb.ChangePwdRes

	serviceAddress ,err := m.FindService(m.ServicName)
	if err!=nil{
		return resp , err
	}

	conn , err := m.CreateConn(serviceAddress)
	if err!=nil{
		return resp , err
	}
	defer conn.Close()

	memberServiceClient := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err = memberServiceClient.ChangePwd(ctx , req)
	if err!=nil{
		return resp , err
	}else {
		return resp , nil
	}
}