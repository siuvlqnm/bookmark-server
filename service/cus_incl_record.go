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

func DeleteInclRecord(i model.CusInclRecord) (err error) {
	err = global.GVA_DB.Delete(&i).Error
	return
}

func GetInclRecordByPageID(SharePageID uint32) (err error, list interface{}) {
	var i []model.CusInclRecord
	err = global.GVA_DB.Where("share_page_id = ?", SharePageID).Order("id DESC").Find(&i).Error
	return err, i
}
