package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusBookmarkRouter(Router *gin.RouterGroup) {
	BookmarkRouter := Router.Group("bookmark").Use(middleware.OperationRecord())
	{
		BookmarkRouter.POST("new", cus.CreateBookmark)
		BookmarkRouter.POST("list", cus.GetBookmarkList)
		BookmarkRouter.PUT("update", cus.UpdateBookmark)
		BookmarkRouter.DELETE("delete", cus.DeleteBookmark)
		BookmarkRouter.PUT("updateToStar", cus.UpdateToStar)
	}
}
