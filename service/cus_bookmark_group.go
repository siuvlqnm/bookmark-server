package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
	"reflect"
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
	err = db.Order("sort ASC").Find(&allGroup).Error
	list = utils.GenerateTree(model.CusBookmarkGroups.ConvertToINodeArray(allGroup), nil)
	return
}

func GetBookmarkGroup(GSeaEngineId uint32, userId uint) (err error, list interface{}) {
	var allGroup []model.CusBookmarkGroup
	var g []model.CusBookmarkGroup
	err = global.GVA_DB.Preload("Bookmark").Where("g_sea_engine_id = ? AND cus_user_id = ?", GSeaEngineId, userId).First(&g).Error
	for i := 0; i < len(g); i++ {
		g[0].GroupParentID = 0
	}
	err = global.GVA_DB.Preload("Bookmark").Where("cus_user_id = ?", userId).Order("sort ASC").Find(&allGroup).Error
	respNodes := utils.FindRelationNode(model.CusBookmarkGroups.ConvertToINodeArray(g), model.CusBookmarkGroups.ConvertToINodeArray(allGroup))
	list = utils.GenerateTree(respNodes, nil)
	return err, list
}

func CreateBookmarkGroup(group model.CusBookmarkGroup) (err error, g *model.CusBookmarkGroup) {
	err = global.GVA_DB.Create(&group).Error
	return err, &group
}

func UpateGroupGSeaEngineId(id uint, GSeaEngineId uint32) {
	var group model.CusBookmarkGroup
	global.GVA_DB.Model(&group).Where("id = ?", id).Update("g_sea_engine_id", GSeaEngineId)
	return
}

func UpdateBookmarkGroup(u *model.CusBookmarkGroup) (err error) {
	var g *model.CusBookmarkGroup
	upDateMap := make(map[string]interface{})
	upDateMap["group_parent_id"] = u.GroupParentID
	upDateMap["group_name"] = u.GroupName
	upDateMap["group_icon"] = u.GroupIcon
	upDateMap["is_archive"] = u.IsArchive
	upDateMap["sort"] = u.Sort
	err = global.GVA_DB.Model(&g).Where("g_sea_engine_id = ?", u.GSeaEngineID).Updates(upDateMap).Error
	return
}

func DeleteBookmarkGroup(GSeaEngineId uint32) (err error) {
	var g model.CusBookmarkGroup
	return global.GVA_DB.Where("g_sea_engine_id = ?", GSeaEngineId).Delete(&g).Error
}

//func GetGroupIdByGSeaEngineId(GSeaEngineId uint32) (groupId int) {
//	val, err := utils.GetSetValue("group", GSeaEngineId)
//	if err != nil {
//		var g model.CusBookmarkGroup
//		err := global.GVA_DB.Select("id").Where("g_sea_engine_id = ?", GSeaEngineId).First(&g).Error
//		if err != nil {
//			return 0
//		}
//		utils.SetSetValue("group", GSeaEngineId, int(g.ID))
//		return int(g.ID)
//	}
//	id, _ := strconv.Atoi(val)
//	return id
//}

func SetBookmarkGroupSort(userId uint, s request.SetGroupSort) (err error) {
	var g model.CusBookmarkGroup
	if s.X-s.Y > 0 {
		err = global.GVA_DB.Model(&g).Where("sort >= ? AND sort < ? AND group_parent_id = ? AND cus_user_id = ?", s.Y, s.X, s.F, userId).UpdateColumn("sort", gorm.Expr("sort + ?", 1)).Error
	} else {
		err = global.GVA_DB.Model(&g).Where("sort > ? AND sort <= ? AND group_parent_id = ? AND cus_user_id = ?", s.X, s.Y, s.F, userId).UpdateColumn("sort", gorm.Expr("sort - ?", 1)).Error
	}
	err = global.GVA_DB.Model(&g).Where("g_sea_engine_id = ? AND cus_user_id = ?", s.G, userId).Update("sort", s.Y).Error
	return
}

func GetBookmarkGroupSort(userId uint, g model.CusBookmarkGroup) (sort int) {
	global.GVA_DB.Select("sort").Debug().Where("group_parent_id = ? AND cus_user_id = ?", g.GroupParentID, userId).Order("sort DESC").Take(&g)
	return g.Sort
}

func CopyBookmarkGroup(r request.CopyBookmarkGroupRequest, userId uint) (err error) {
	var allGroup []model.CusBookmarkGroup
	var g []model.CusBookmarkGroup
	global.GVA_DB.Where("g_sea_engine_id = ? AND cus_user_id = ?", r.CusGroupID, userId).First(&g)
	g[0].GroupParentID = 0
	global.GVA_DB.Where("cus_user_id = ?", userId).Order("sort ASC").Find(&allGroup)
	respNodes := utils.FindRelationNode(model.CusBookmarkGroups.ConvertToINodeArray(g), model.CusBookmarkGroups.ConvertToINodeArray(allGroup))

	var sp = make([]model.CusShareGroup, len(respNodes))
	for i, _ := range respNodes {
		val := reflect.ValueOf(respNodes).Index(i).Elem()
		sp[i].CusUserID = uint(val.FieldByName("CusUserID").Uint())
		sp[i].SharePageID = r.PSeaEngineID
		sp[i].SGSeaEngineID = uint32(val.FieldByName("GSeaEngineID").Uint())
		sp[i].GroupParentID = int(val.FieldByName("GroupParentID").Int())
		sp[i].GroupName = val.FieldByName("GroupName").String()
		sp[i].GroupIcon = val.FieldByName("GroupIcon").String()
		if sp[i].GroupParentID == 0 {
			sp[i].Sort = GetShareGroupSort(userId, sp[i].GroupParentID, sp[i].SharePageID) + 1
		} else {
			sp[i].Sort = int(val.FieldByName("Sort").Int())
		}
	}
	err = global.GVA_DB.Create(&sp).Error
	return
}
