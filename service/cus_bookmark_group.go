package service

import (
	"strconv"

	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/utils"
)

func GetBookmarkGroupList(userId uint, where request.GetGetBookmarkGroup) (err error, list interface{}) {
	var allGroup []model.CusBookmarkGroup
	err = global.GVA_DB.Where("cus_user_id = ? AND is_archive = ?", userId, where.IsArchive).Order("id desc").Find(&allGroup).Error
	list = utils.GenerateTree(model.CusBookmarkGroups.ConvertToINodeArray(allGroup), nil)
	return
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
	if val, err := utils.GetSetValue("group", GSeaEngineId); err != nil {
		var g model.CusBookmarkGroup
		err := global.GVA_DB.Select("id").Where("g_sea_engine_id = ?", GSeaEngineId).Find(&g).Error
		if err != nil {
			return 0
		}
		utils.SetSetValue("group", GSeaEngineId, int(g.ID))
		return int(g.ID)
	} else {
		id, _ := strconv.Atoi(val)
		return id
	}
}
