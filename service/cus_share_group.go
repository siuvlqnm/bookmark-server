package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

func GetShareGroupSort(userId uint, GroupParentID int, SharePageID uint32) (sort int) {
	var g model.CusShareGroup
	global.GVA_DB.Select("sort").Where("cus_user_id = ? AND group_parent_id = ? AND share_page_id = ?", userId, GroupParentID, SharePageID).Order("sort DESC").Take(&g)
	return g.Sort
}

func CreateShareGroup(g model.CusShareGroup) (err error) {
	err = global.GVA_DB.Create(&g).Error
	return
}

func UpdateShareGroupSGSeaEngineID(id uint, SGSeaEngineID uint32) {
	var g model.CusShareGroup
	global.GVA_DB.Model(&g).Where("id = ?", id).Update("sg_sea_engine_id", SGSeaEngineID)
	return
}

func UpdateShareGroup(userId uint, g model.CusShareGroup) (err error) {
	db := global.GVA_DB.Model(&g).Where("cus_user_id = ? AND sg_sea_engine_id = ?", userId, g.SGSeaEngineID)
	if g.GroupName != "" {
		db = db.Update("group_name", g.GroupName)
	}
	err = db.Update("group_icon", g.GroupIcon).Update("group_parent_id", g.GroupParentID).Update("share_page_id", g.SharePageID).Error
	return
}

func DeleteShareGroup(SGSeaEngineID uint32) (err error) {
	var s model.CusShareGroup
	err = global.GVA_DB.Where("sg_sea_engine_id = ?", SGSeaEngineID).Delete(&s).Error
	return err
}

func SetShareGroupSort(userId uint, s request.SetShareGroupSort) (err error) {
	var g model.CusShareGroup
	if s.X-s.Y > 0 {
		err = global.GVA_DB.Model(&g).Where("sort >= ? AND sort < ? AND group_parent_id = ? AND cus_user_id = ?", s.Y, s.X, s.F, userId).UpdateColumn("sort", gorm.Expr("sort + ?", 1)).Error
	} else {
		err = global.GVA_DB.Model(&g).Where("sort > ? AND sort <= ? AND group_parent_id = ? AND cus_user_id = ?", s.X, s.Y, s.F, userId).UpdateColumn("sort", gorm.Expr("sort - ?", 1)).Error
	}
	err = global.GVA_DB.Model(&g).Where("sg_sea_engine_id = ? AND cus_user_id = ?", s.G, userId).Update("sort", s.Y).Error
	return
}

func GetShareGroup(SharePageID uint32, userId uint) (err error, list interface{}) {
	var g []model.CusShareGroup
	err = global.GVA_DB.Preload("Bookmark").Where("share_page_id = ? AND cus_user_id = ?", SharePageID, userId).Order("sort ASC").Find(&g).Error
	list = utils.GenerateTree(model.CusShareGroups.ConvertToINodeArray(g), nil)
	return err, list
}
