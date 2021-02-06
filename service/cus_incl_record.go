package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func CreateInclRecord(i model.CusInclRecord) (err error) {
	err = global.GVA_DB.Create(&i).Error
	return
}

func UpdateInclRecord(i model.CusInclRecord) (err error) {
	err = global.GVA_DB.Updates(&i).Error
	return
}
