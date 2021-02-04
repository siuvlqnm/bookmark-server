package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusBookmarkGroupRouter(Router *gin.RouterGroup) {
	BookmarkGroupRouer := Router.Group("folder").Use(middleware.OperationRecord())
	{
		BookmarkGroupRouer.GET("all", cus.GetAllBookmarkGroup)
		BookmarkGroupRouer.POST("new", cus.CreateNewGroup)
		BookmarkGroupRouer.PUT("update", cus.UpdateBookmarkGroup)
		BookmarkGroupRouer.DELETE("delete", cus.DeleteBookmarkGroup)
		BookmarkGroupRouer.GET("/get/:gSeaEngineId", cus.GetBookmarkGroupWithBookmark)
		BookmarkGroupRouer.PUT("sort", cus.SetBookmarkGroupSort)
		BookmarkGroupRouer.POST("copyGroup", cus.CopyBookmarkGroup)
	}
}
