package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
	"go.uber.org/zap"
)

func GetAllBookmarkGroup(c *gin.Context) {
	var g request.GetGetBookmarkGroup
	_ = c.ShouldBindQuery(&g)
	err, list := service.GetAllBookmarkGroup(getUserID(c), g)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func GetBookmarkGroupWithBookmark(c *gin.Context) {
	var g request.GetGetBookmarkGroup
	_ = c.ShouldBindUri(&g)
	err, list := service.GetBookmarkGroup(g.GSeaEngineId, getUserID(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func CreateNewGroup(c *gin.Context) {
	var group model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&group)
	if err := utils.Verify(group, utils.NewBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := getUserID(c)
	group.CusUserID = userId
	sort := service.GetBookmarkGroupSort(userId, group)
	group.Sort = sort + 1
	err, cbg := service.CreateBookmarkGroup(group)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	}
	murmur32 := utils.GetMurmur32("group:", cbg.ID)
	service.UpateGroupGSeaEngineId(cbg.ID, murmur32)
	response.OkWithMessage("添加成功", c)
}

func UpdateBookmarkGroup(c *gin.Context) {
	var u model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&u)
	if err := utils.Verify(u, utils.UpdateBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.UpdateBookmarkGroup(&u)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func DeleteBookmarkGroup(c *gin.Context) {
	var d model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&d)
	if err := utils.Verify(d, utils.UpdateBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.DeleteBookmarkGroup(d.GSeaEngineID)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func SetBookmarkGroupSort(c *gin.Context) {
	var s request.SetGroupSort
	_ = c.ShouldBindJSON(&s)
	err := service.SetBookmarkGroupSort(getUserID(c), s)
	if err != nil {
		global.GVA_LOG.Error("排序失败", zap.Any("err", err))
		response.FailWithMessage("排序失败", c)
		return
	}
	response.OkWithMessage("排序成功", c)
}

func CopyBookmarkGroup(c *gin.Context) {
	var r request.CopyBookmarkGroupRequest
	_ = c.ShouldBindJSON(&r)

	err := service.CopyBookmarkGroup(r, getUserID(c))
	if err != nil {
		global.GVA_LOG.Error("复制失败", zap.Any("err", err))
		response.FailWithMessage("复制失败", c)
	}
	response.OkWithMessage("复制成功", c)
}
