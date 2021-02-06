package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
	"go.uber.org/zap"
)

func CreateInclRecord(c *gin.Context) {
	var i model.CusInclRecord
	_ = c.ShouldBindJSON(&i)
	if err := utils.Verify(i, utils.CreateInclRecordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	i.InclUserID = getUserID(c)
	err := service.CreateInclRecord(i)
	if err != nil {
		global.GVA_LOG.Error("提交失败", zap.Any("err", err))
		response.FailWithMessage("提交失败", c)
		return
	}
	response.OkWithMessage("提交成功", c)
}

func UpdateInclRecord(c *gin.Context) {

}

func DeleteInclRecord(c *gin.Context) {
	var i model.CusInclRecord

	err := service.DeleteInclRecord(i)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func AuditInclRecord(c *gin.Context) {
	var i model.CusInclRecord
	_ = c.ShouldBindJSON(&i)

	err := service.UpdateInclRecord(i)
	if err != nil {
		global.GVA_LOG.Error("审核失败", zap.Any("err", err))
		response.FailWithMessage("审核失败", c)
		return
	}
	if i.IsAccept == 1 {
		_, w := service.GetWebSite(i.Domain)
		b := model.CusBookmark{CusWebID: w.ID, CusUserID: service.GetSharePageUserIDByPSeaEngineID(i.SharePageID), TargetUrl: i.TargetUrl, Domain: i.Domain, Path: i.Path, Query: i.Query, Title: i.Title, Description: i.Description, ShareGroupID: i.ShareGroupID}
		if err, cbm := service.CreateBookmark(b); err != nil {
			global.GVA_LOG.Error("审核失败", zap.Any("err", err))
			response.FailWithMessage("审核失败", c)
			return
		} else {
			murmur32 := utils.GetMurmur32("bookmark:", cbm.ID)
			service.UpdateBookmarkMSeaEngineId(int(cbm.ID), murmur32)
			response.OkWithMessage("审核成功", c)
			return
		}
	}
}
