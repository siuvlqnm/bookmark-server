package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/model"
)

func CreateShareGroup(c *gin.Context) {
	var s model.CusShareGroup
	_ = c.ShouldBindJSON(&s)

}

func UpdateShareGroup(c *gin.Context) {

}

func DeleteShareGroup(c *gin.Context) {

}
