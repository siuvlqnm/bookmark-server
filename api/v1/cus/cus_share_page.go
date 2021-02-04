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

func GetSharePageList(c *gin.Context) {
	err, list := service.GetSharePageList(getUserID(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func CreateSharePage(c *gin.Context) {
	var s model.CusSharePage
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.CreateSharePageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := getUserID(c)
	s.CusUserID = userId
	sort := service.GetSharePageSort(userId, s)
	s.Sort = sort + 1
	if s.IsPassword {
		s.PagePassword = utils.MD5V([]byte(s.PagePassword))
	}
	err, csp := service.CreateSharePage(s)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	}
	murmur32 := utils.GetMurmur32("sharePage:", csp.ID)
	service.UpatePagePSeaEngineId(int(csp.ID), murmur32)
	response.OkWithMessage("添加成功", c)
}

func UpdateSharePage(c *gin.Context) {
	var s model.CusSharePage
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.CreateSharePageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if s.IsPassword {
		s.PagePassword = utils.MD5V([]byte(s.PagePassword))
	}
	err := service.UpdateSharePage(getUserID(c), s)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func DeleteSharePage(c *gin.Context) {
	var s model.CusSharePage
	_ = c.ShouldBindJSON(&s)

	err := service.DeleteSharePage(s.PSeaEngineID)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func SetSharePageSort(c *gin.Context) {
	var p request.SetPageSort
	_ = c.ShouldBindJSON(&p)
	err := service.SetSharePageSort(getUserID(c), p)
	if err != nil {
		global.GVA_LOG.Error("排序失败", zap.Any("err", err))
		response.FailWithMessage("排序失败", c)
		return
	}
	response.OkWithMessage("排序成功", c)
}
