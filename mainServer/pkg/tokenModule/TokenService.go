package tokenModule

import (
	// "fmt"
	"time"
	"context"
	pb "mainServer/proto"
	"mainServer/pkg/base"
)

type TokenService struct{
	base.Service
	ServicName  string 
}

func NewTokenService() *TokenService{
	var tokenService TokenService
	tokenService.ServicName = "memberService"
	return &tokenService
}

func (t *TokenService) IsValidJwt(req *pb.JwtToken)(*pb.IsValidJwtRes,error) {
	var resp  *pb.IsValidJwtRes

	serviceAddress ,err := t.FindService(t.ServicName)
	if err!=nil{
		return resp , err
	}

	conn , err := t.CreateConn(serviceAddress)
	if err!=nil{
		return resp , err
	}
	defer conn.Close()

	memberServiceClient := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err = memberServiceClient.IsValidJwt(ctx , req)
	if err!=nil{
		return resp , err
	}else {
		return resp , nil
	}
}

func (t *TokenService) IsJwtInTime(req *pb.JwtToken)(*pb.IsJwtInTimeRes,error){
	
	var resp  *pb.IsJwtInTimeRes

	serviceAddress ,err := t.FindService(t.ServicName)
	if err!=nil{
		return resp , err
	}

	conn , err := t.CreateConn(serviceAddress)
	if err!=nil{
		return resp , err
	}
	defer conn.Close()

	memberServiceClient := pb.NewMemberServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err = memberServiceClient.IsJwtInTime(ctx , req)
	if err!=nil{
		return resp , err
	}else {
		return resp , nil
	}
}