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

func GetBookmarkGroupList(c *gin.Context) {
	var g request.GetGetBookmarkGroup
	_ = c.ShouldBindQuery(&g)
	if err, list := service.GetBookmarkGroupList(getUserID(c), g); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

func CreateNewGroup(c *gin.Context) {
	var group model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&group)
	if err := utils.Verify(group, utils.NewBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	group.CusUserId = getUserID(c)
	if err, cbg := service.CreateBookmarkGroup(group); err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	} else {
		murmur32 := utils.GetMurmur32("group:", int(cbg.ID))
		service.UpateGroupGSeaEngineId(int(cbg.ID), murmur32)
		response.OkWithMessage("添加成功", c)
	}
}

func UpdateBookmarkGroup(c *gin.Context) {
	var u model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&u)
	if err := utils.Verify(u, utils.UpdateBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateBookmarkGroup(&u); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
		return
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func DeleteBookmarkGroup(c *gin.Context) {
	var d model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&d)
	if err := utils.Verify(d, utils.UpdateBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBookmarkGroup(d.GSeaEngineId); err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
