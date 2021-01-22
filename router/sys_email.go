package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/siuvlqnm/bookmark/api/v1"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
