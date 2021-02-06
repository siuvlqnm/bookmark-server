package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
)

func InitCusInclRecordRouter(Router *gin.RouterGroup) {
	InclRecordRouter := Router.Group("incl")
	{
		InclRecordRouter.GET("get", cus.GetSharePageInclRecord)
		InclRecordRouter.DELETE("delete", cus.DeleteInclRecord)
		InclRecordRouter.PUT("audit", cus.AuditInclRecord)
	}
}

func InitPublicCusInclRecordRouter(Router *gin.RouterGroup) {
	PublicInclRecordRouter := Router.Group("incl")
	{
		PublicInclRecordRouter.POST("new", cus.CreateInclRecord)
	}
}
