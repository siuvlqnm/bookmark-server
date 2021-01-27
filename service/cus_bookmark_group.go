package service

import (
	"strconv"

	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/utils"
)

func GetAllBookmarkGroup(userId uint, where request.GetGetBookmarkGroup) (err error, list interface{}) {
	var allGroup []model.CusBookmarkGroup
	var g model.CusBookmarkGroup
	db := global.GVA_DB.Model(&g).Where("cus_user_id = ?", userId)
	if where.IsArchive {
		db = db.Where("is_archive = ?", true)
	} else {
		db = db.Where("is_archive = ?", false)
	}
	err = db.Order("id desc").Find(&allGroup).Error
	list = utils.GenerateTree(model.CusBookmarkGroups.ConvertToINodeArray(allGroup), nil)
	return
}

func GetBookmarkGroup(GSeaEngineId uint32) (err error, list interface{}) {
	var allGroup []model.CusBookmarkGroup
	var g []model.CusBookmarkGroup
	err = global.GVA_DB.Preload("Bookmark").Where("g_sea_engine_id = ?", GSeaEngineId).First(&g).Error
	for i := 0; i < len(g); i++ {
		g[0].GroupParentId = 0
	}
	err = global.GVA_DB.Preload("Bookmark").Find(&allGroup).Error
	respNodes := utils.FindRelationNode(model.CusBookmarkGroups.ConvertToINodeArray(g), model.CusBookmarkGroups.ConvertToINodeArray(allGroup))
	list = utils.GenerateTree(respNodes, nil)
	return err, list
}

func CreateBookmarkGroup(group model.CusBookmarkGroup) (err error, g *model.CusBookmarkGroup) {
	err = global.GVA_DB.Create(&group).Error
	return err, &group
}

func UpateGroupGSeaEngineId(id int, GSeaEngineId uint32) (err error) {
	var group model.CusBookmarkGroup
	err = global.GVA_DB.Model(&group).Where("id = ?", id).Update("g_sea_engine_id", GSeaEngineId).Error
	return
}

func UpdateBookmarkGroup(u *model.CusBookmarkGroup) (err error) {
	var g *model.CusBookmarkGroup
	upDateMap := make(map[string]interface{})
	upDateMap["group_parent_id"] = u.GroupParentId
	upDateMap["group_name"] = u.GroupName
	upDateMap["group_icon"] = u.GroupIcon
	upDateMap["is_archive"] = u.IsArchive
	err = global.GVA_DB.Model(&g).Where("g_sea_engine_id = ?", u.GSeaEngineId).Updates(upDateMap).Error
	return
}

func DeleteBookmarkGroup(GSeaEngineId uint32) (err error) {
	var g model.CusBookmarkGroup
	return global.GVA_DB.Where("g_sea_engine_id = ?", GSeaEngineId).Delete(&g).Error
}

func GetGroupIdByGSeaEngineId(GSeaEngineId uint32) (groupId int) {
	val, err := utils.GetSetValue("group", GSeaEngineId)
	if err != nil {
		var g model.CusBookmarkGroup
		err := global.GVA_DB.Select("id").Where("g_sea_engine_id = ?", GSeaEngineId).First(&g).Error
		if err != nil {
			return 0
		}
		utils.SetSetValue("group", GSeaEngineId, int(g.ID))
		return int(g.ID)
	}
	id, _ := strconv.Atoi(val)
	return id
}
