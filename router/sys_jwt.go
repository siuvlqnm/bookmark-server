package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/siuvlqnm/bookmark/api/v1"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
