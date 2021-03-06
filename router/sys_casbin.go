package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/siuvlqnm/bookmark/api/v1"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
