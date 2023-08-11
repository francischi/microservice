package main

import (
	"fmt"
	"net/http"
	"mainServer/pkg"
	"mainServer/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func main(){

	if err := helpers.InitEnvSetting();err !=nil{
		panic(err)
	}
	router := gin.New()
	router.Use(gin.Logger())
	pkg.SetRouter(router)
	router.GET("/" , func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"success" : true, 
			"message" : "service is running", 
		})
	})

	port := helpers.GetEnvStr("PORT")
	serverPort := fmt.Sprintf(":%s",port)
	router.Run(serverPort)
}