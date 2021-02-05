package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
)

func InitCusMaintRecordRouter(Router *gin.RouterGroup) {
	MaintRecordRouter := Router.Group("maint")
	{
		MaintRecordRouter.GET("get")
		MaintRecordRouter.POST("new", cus.CreateMaintRecord)
		MaintRecordRouter.PUT("update", cus.UpdateMaintRecord)
		MaintRecordRouter.DELETE("delete", cus.DeleteMaintRecord)
	}
}

func InitPublicCusMaintRecordRouter() {

}
