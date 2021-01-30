package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func GetSharePageSort(userId uint, s model.CusSharePage) (sort int) {
	global.GVA_DB.Select("sort").Where("cus_user_id = ?", userId).Order("sort DESC").Take(&s)
	return s.Sort
}

func CreateSharePage(s model.CusSharePage) (err error, page model.CusSharePage) {
	err = global.GVA_DB.Create(&s).Error
	return err, s
}

func UpdateSharePage(userId uint, s model.CusSharePage) (err error) {
	db := global.GVA_DB.Model(&s).Where("cus_user_id = ? AND p_sea_engine_id = ?", userId, s.PSeaEngineID)
	if s.PageName != "" {
		db = db.Update("page_name", s.PageName)
	}
	err = db.Update("is_password", s.IsPassword).Update("page_password", s.PagePassword).Error
	return
}

func UpatePagePSeaEngineId(id int, PSeaEngineID uint32) (err error) {
	var s model.CusSharePage
	err = global.GVA_DB.Model(&s).Where("id = ?", id).Update("p_sea_engine_id", PSeaEngineID).Error
	return
}

func DeleteSharePage(PSeaEngineID uint32) (err error) {
	var s model.CusSharePage
	err = global.GVA_DB.Where("p_sea_engine_id = ?", PSeaEngineID).Delete(&s).Error
	return
}

func GetSharePageList(userId uint) (err error, list interface{}) {
	var all []model.CusSharePage
	err = global.GVA_DB.Where("cus_user_id = ?", userId).Order("sort ASC").Find(&all).Error
	return err, all
}