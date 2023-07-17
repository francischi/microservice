package base

import (
	"mainServer/pkg/except"
	"github.com/gin-gonic/gin"
)

type MiddleWare struct{
}

func (m *MiddleWare)InvaliAugument(g *gin.Context , errorMsg string){
	m.HandleError(g , &except.InvalidArgument{Message: errorMsg})
}

func (m *MiddleWare)SystemError(g *gin.Context , errorMsg string){
	m.HandleError(g , &except.SystemError{Message: errorMsg})
}

func (m *MiddleWare)HandleError(g *gin.Context , err error){
	statusCode,errMsg := except.Error2StatusCode(err)
	g.JSON(statusCode,gin.H{
		"success":false,
		"error_msg":errMsg,
		"result":"",
	})
	g.Abort()
	return 
}