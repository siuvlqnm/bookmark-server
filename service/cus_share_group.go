package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func GetShareGroupSort(userId uint, g model.CusShareGroup) (sort int) {
	global.GVA_DB.Select("sort").Where("cus_user_id = ? AND group_parent_id = ?", userId, g.GroupParentID).Order("sort DESC").Take(&g)
	return g.Sort
}

func CreateShareGroup(g model.CusShareGroup) (err error, sg *model.CusShareGroup) {
	err = global.GVA_DB.Create(&g).Error
	return err, &g
}

func UpdateShareGroupSGSeaEngineID(id uint, SGSeaEngineID uint32) {
	var g model.CusShareGroup
	global.GVA_DB.Model(&g).Where("id = ?", id).Update("s_g_sea_engine_id", SGSeaEngineID)
	return
}

func UpdateShareGroup(userId uint, g model.CusShareGroup) (err error) {
	db := global.GVA_DB.Model(&g).Where("cus_user_id = ? AND s_g_sea_engine_id = ?", userId, g.SGSeaEngineID)
	if g.GroupName != "" {
		db = db.Update("group_name", g.GroupName)
	}
	err = db.Update("group_icon", g.GroupIcon).Update("group_parent_id", g.GroupParentID).Update("share_page_id", g.SharePageID).Error
	return
}

func DeleteShareGroup(SGSeaEngineID uint32) (err error) {
	var s model.CusShareGroup
	err = global.GVA_DB.Where("s_g_sea_engine_id = ?", SGSeaEngineID).Delete(&s).Error
	return err
}
