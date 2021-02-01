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

func CreateShareGroup(c *gin.Context) {
	var g model.CusShareGroup
	_ = c.ShouldBindJSON(&g)
	if err := utils.Verify(g, utils.CreateShareGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := getUserID(c)
	sort := service.GetShareGroupSort(userId, g)
	g.CusUserID = userId
	g.Sort = sort + 1
	err, csg := service.CreateShareGroup(g)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	}
	murmur32 := utils.GetMurmur32("shareGroup:", csg.ID)
	service.UpdateShareGroupSGSeaEngineID(csg.ID, murmur32)
	response.OkWithMessage("添加成功", c)
}

func UpdateShareGroup(c *gin.Context) {
	var g model.CusShareGroup
	_ = c.ShouldBindJSON(&g)
	if err := utils.Verify(g, utils.CreateShareGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.UpdateShareGroup(getUserID(c), g)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func DeleteShareGroup(c *gin.Context) {
	var s model.CusShareGroup
	_ = c.ShouldBindJSON(&s)

	err := service.DeleteShareGroup(s.SGSeaEngineID)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func SetShareGroupSort(c *gin.Context) {
	var s request.SetShareGroupSort
	_ = c.ShouldBindJSON(&s)
	err := service.SetShareGroupSort(getUserID(c), s)
	if err != nil {
		global.GVA_LOG.Error("排序失败", zap.Any("err", err))
		response.FailWithMessage("排序失败", c)
		return
	}
	response.OkWithMessage("排序成功", c)
}
