package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
)

func InitCusShareGroupRouter(Router *gin.RouterGroup) {
	ShareGroupRouter := Router.Group("")
	{
		ShareGroupRouter.POST("new", cus.CreateShareGroup)
		ShareGroupRouter.PUT("update", cus.UpdateShareGroup)
		ShareGroupRouter.DELETE("delete", cus.DeleteShareGroup)
	}
}
