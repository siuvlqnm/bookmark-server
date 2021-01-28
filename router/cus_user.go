package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.PUT("changePassword", cus.ChangePassword)
	}
}
