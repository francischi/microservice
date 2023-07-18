# **microservice (golang)**

## **Introduction**
```
本專案將以會員系統示範使用golang建置並以clean architecture為基礎之微服務架構。
```

## 技術點
* clean architecture
* microservice
* docker
* resful api
* grpc
* consul


## **Microservice**
```
將原本單體式架構拆成各個小服務，能夠增加服務開發與部屬上的獨立性，專案擴大時開發管理的成本也能有效的下降
```
**優點**
* 擴展容易 ( Auto-Scaling )
* 各服務間解耦
* 縮短編譯時間
* 開發不受語言限制
* 故障影響範圍限縮

<br>


## **Architecture**
```
restful 較能展現對於資源的使用意圖，且串接較容易因此對於外部請求使用restful api進行溝通，
透過mainServer進行restful api對對應服務的grpc request轉發，
各個服務之間則使用gRPC進行溝通，使用註冊發現中心能夠更好的管理服務且能夠進行服務的擴展與平衡(load balance)。
```

* 對外 : restful API
* 服務間 : gRPC
* 服務間管理 : 註冊發現中心 (consul)

### mainServer
```
主要server進行restful api請求處理並進行對內部服務之grpc request封裝及轉發，透過對註冊發現中心(consul)請求特定服務之地址
完成服務主機的定位與grpc request轉發。
```
### server
```
會依據不同功能domain拆分，負責接收grpc request，並實作商業邏輯，以clean architecture進行建置，能使商業邏輯與服務解耦，
且能夠增加架構調整與編寫測時時的商業邏輯程式碼健壯性。
```
### consul
```
註冊發現中心，提供server啟動時進行註冊，能夠增加服務主機的管理與分流之便利性。
```

<br>


## **note**
載golang grpc套件
```
go get github.com/golang/protobuf/protoc-gen-go 
go get -t github.com/grpc/grpc-go    
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

生成protocal buffer
```
protoc --go-grpc_out=. ./proto/*.proto
protoc --go_out=. ./proto/*.proto
```

安裝註冊中心
```
docker pull hashicorp/consul:latest
docker run -d -p 8500:8500 --restart=always --name=consul hashicorp/consul:latest agent -server -bootstrap -ui -node=1 -client=0.0.0.0
```