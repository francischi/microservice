package base

import (
	// "fmt"
	// "errors"
	// "github.com/gin-gonic/gin"
	"server/pkg/except"
	"google.golang.org/grpc/codes"
)

type Controller struct {
}

// func (c *Controller) SuccessRes(g *gin.Context ,result interface{}) {
// 	g.JSON(200,gin.H{
// 		"success":true,
// 		"error_msg":nil,
// 		"result":result,
// 	})
// }

// func (c *Controller)HandleError(g *gin.Context , err error){
// 	statusCode := except.Error2StatusCode(err)
// 	g.JSON(statusCode,gin.H{
// 		"success":false,
// 		"error_msg":err.Error(),
// 		"result":nil,
// 	})
// }

func (c *Controller)Error2StatusCode(err error)(codes.Code){

	statusCode := except.Error2StatusCode(err)
	
	return statusCode
}