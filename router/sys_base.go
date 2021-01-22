package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/siuvlqnm/bookmark/api/v1"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.POST("/c/register", cus.Register)
		BaseRouter.POST("/c/login", cus.Login)
	}
	return BaseRouter
}
