package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusWebsiteRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("website").Use(middleware.OperationRecord())
	{
		UserRouter.POST("webInfo", cus.GetWebInfo)
	}
}
