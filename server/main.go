package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"strconv"
	"math/rand"
	"encoding/hex"
	pb "server/proto"
	"server/pkg/helpers"
	"google.golang.org/grpc"
	"server/pkg"
	// "server/pkg/memberModule"
	"github.com/hashicorp/consul/api"
)

const (
	ttl = 30 * time.Second
)

func RegisterToConsul(registrationAddress string ,serviceName string, serviceHost string ,servicePort int){
	config := api.Config{
		Address: registrationAddress,
	}
	registry, err := api.NewClient(&config)
	if err != nil {
		log.Fatalln(err)
	}
	var h [16]byte
	rand.Read(h[:])
	// 生成一个全局ID
	id := fmt.Sprintf("helloserver-%s-%d", hex.EncodeToString(h[:]), servicePort)
	fmt.Println(id)
	// 註冊到 Consul，包含地址、端口訊息，以及健康檢查
	err = registry.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
		Name:    serviceName,
		Port:    servicePort,
		Address: serviceHost,
		Check: &api.AgentServiceCheck{
			TTL:     (ttl + time.Second).String(),
			Timeout: time.Minute.String(),
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		checkid := "service:" + id
		for range time.Tick(ttl) {
			err := registry.Agent().PassTTL(checkid, "")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}

func main() {
	if err := helpers.InitEnvSetting();err!=nil{
		log.Fatalf("init env err: %v", err)
	}
	serviceHost := helpers.GetEnvStr("HOST")
	servicePort := helpers.GetEnvStr("PORT")
	serviceName := helpers.GetEnvStr("SERVICE.NAME")
	rcAddress := helpers.GetEnvStr("REGISTRATION_CENTER.ADDRESS")

	if err,_ := helpers.InitMySql();err !=nil{
		log.Fatalf("failed to connect db: %v", err)
	}
	// Create gRPC Server
	address := fmt.Sprintf(":%s",servicePort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	servicePortInt, err := strconv.Atoi(servicePort)
	if err!=nil{
		log.Fatalf("env err: %v", err)
	}

	RegisterToConsul(rcAddress ,serviceName , serviceHost , servicePortInt)

	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	
	var memberService pkg.MemberService

	pb.RegisterMemberServiceServer(s, &memberService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
