package base

import (
	"os"
	"fmt"
	"time"
	"errors"
	"strconv"
	"math/rand"
	"mainServer/pkg/except"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

type Service struct{

}

func (b *Service)FindService(serviceName string)(address string ,err error){
	envCfg := b.GetEnv()

	cfg := api.DefaultConfig()
    cfg.Address = envCfg.RcAddress

    client, err := api.NewClient(cfg)
    if err != nil {
        return "", errors.New("cant init register center client")
    }

	datas, _, err := client.Health().Service(serviceName, "", true, nil)

	if len(datas)==0 || err!=nil{
		return "",  errors.New("cant find service")
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(datas))
	selectedData := datas[random]
	selectedPort := selectedData.Service.Port
	selectedHost := selectedData.Service.Address

	address = fmt.Sprintf("%s:%s" , selectedHost ,strconv.Itoa(selectedPort))
	return address ,nil
}

func (b *Service) CreateConn(serviceAddress string)(*grpc.ClientConn ,error){
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithBlock())
	return conn , err
}

func (b *Service) SuccessRes(g *gin.Context ,result interface{}) {
	g.JSON(200,gin.H{
		"success":true,
		"error_msg":nil,
		"result":result,
	})
}

func (b *Service)HandleError(g *gin.Context , err error){
	// statusCode := except.Error2StatusCode(err)
	fmt.Println("-----------------------")
	g.JSON(400,gin.H{
		"success":false,
		"error_msg":err.Error(),
		"result":nil,
	})
}

func(s *Service) InvalidArgument(msg string)error{
	var InvalidArgument except.InvalidArgument
	InvalidArgument.Message = msg
	return &InvalidArgument
}

func(s *Service) SystemError(msg string)error{
	var SystemError except.SystemError
	SystemError.Message = msg
	return &SystemError
}

type envConfig struct {
	RcAddress string
}

func (b *Service) GetEnv() envConfig {
	rcAddress := os.Getenv("registrationCenter.address")

	var  envConfig envConfig
	envConfig.RcAddress = rcAddress
	return envConfig
}