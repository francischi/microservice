```
go get github.com/golang/protobuf/protoc-gen-go 
go get -t github.com/grpc/grpc-go    
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

生成protocal buffer
```
protoc --go-grpc_out=. ./proto/ *.proto
protoc --go_out=. ./proto/*.proto
```

安裝註冊中心
```
docker pull hashicorp/consul:latest
docker run -d -p 8500:8500 --restart=always --name=consul hashicorp/consul:latest agent -server -bootstrap -ui -node=1 -client=0.0.0.0
```