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
		if err := service.CreateBookmark(b); err != nil {
			global.GVA_LOG.Error("审核失败", zap.Any("err", err))
			response.FailWithMessage("审核失败", c)
			return
		}
		response.OkWithMessage("审核成功", c)
	}
}

func GetSharePageInclRecord(c *gin.Context) {
	var i model.CusInclRecord
	_ = c.ShouldBindJSON(&i)
	err, list := service.GetInclRecordByPageID(i.SharePageID)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}
