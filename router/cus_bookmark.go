package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusBookmarkRouter(Router *gin.RouterGroup) {
	BookmarkRouter := Router.Group("c").Use(middleware.OperationRecord())
	{
		BookmarkRouter.POST("/bookmark/new", cus.CreateBookmark)
		BookmarkRouter.POST("/bookmark/list", cus.GetBookmarkList)
		BookmarkRouter.PUT("/bookmark/update", cus.UpdateBookmark)
		BookmarkRouter.DELETE("/bookmark/delete", cus.DeleteBookmark)
		BookmarkRouter.PUT("/bookmark/updateToStar", cus.UpdateToStar)
	}
}
