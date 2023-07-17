package tokenModule

import (
	// "fmt"
	"context"
   "google.golang.org/grpc/status"
   "server/pkg/base"
   pb "server/proto"
)

type TokenController struct{
	TokenService TokenService
	// pb.UnimplementedMemberServiceServer
	base.Controller
}

func NewTokenController(tokenService *TokenService) *TokenController{
	return &TokenController{
		TokenService : *tokenService,
	}
}

func (c *TokenController) IsValidJwt(ctx context.Context , req *pb.JwtToken)(*pb.IsValidJwtRes , error){
	ok , err := c.TokenService.IsValidJwt(req.Token)
	if err!=nil {
		return &pb.IsValidJwtRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	if !ok {
		return &pb.IsValidJwtRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	return &pb.IsValidJwtRes{Success: true , ErrMsg: ""} ,nil
}

func (c *TokenController) IsJwtInTime(ctx context.Context , req *pb.JwtToken)(*pb.IsJwtInTimeRes , error){
	ok , err := c.TokenService.IsJwtInTime(req.Token)
	if err!=nil {
		return &pb.IsJwtInTimeRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	if !ok {
		return &pb.IsJwtInTimeRes{
			Success: false , 
			ErrMsg: err.Error(),
		} ,
		status.Error( 
			c.Error2StatusCode(err) , 
			err.Error(),
		)
	}
	return &pb.IsJwtInTimeRes{Success: true , ErrMsg: ""} ,nil
}