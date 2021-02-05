package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/utils"
)

func CreateMaintRecord(c *gin.Context) {
	var m model.CusMaintRecord
	_ = c.ShouldBindJSON(&m)
	if err := utils.Verify(m, utils.ApiVerify); err != nil {

	}
}

func UpdateMaintRecord(c *gin.Context) {

}

func DeleteMaintRecord(c *gin.Context) {

}
