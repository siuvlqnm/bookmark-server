package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
)

func InitCusSharePageRouter(Router *gin.RouterGroup) {
	SharePageRouter := Router.Group("page")
	{
		SharePageRouter.GET("all", cus.GetSharePageList)
		SharePageRouter.POST("new", cus.CreateSharePage)
		SharePageRouter.PUT("update", cus.UpdateSharePage)
		SharePageRouter.DELETE("delete", cus.DeleteSharePage)
		SharePageRouter.PUT("sort")
	}
}

func InitPublicCusSharePageRouter(Router *gin.RouterGroup) {
	PublicSharePageRouter := Router.Group("page")
	{
		PublicSharePageRouter.GET("/get/:pSeaEngineId")
	}
}
