package service

import (
	"fmt"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
	"reflect"
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

func UpatePagePSeaEngineId(id int, PSeaEngineID uint32) {
	var s model.CusSharePage
	global.GVA_DB.Model(&s).Where("id = ?", id).Update("p_sea_engine_id", PSeaEngineID)
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

func SetSharePageSort(userId uint, s request.SetPageSort) (err error) {
	var p model.CusSharePage
	if s.X-s.Y > 0 {
		err = global.GVA_DB.Model(&p).Where("sort >= ? AND sort < ? AND group_parent_id = ? AND cus_user_id = ?", s.Y, s.X, s.F, userId).UpdateColumn("sort", gorm.Expr("sort + ?", 1)).Error
	} else {
		err = global.GVA_DB.Model(&p).Where("sort > ? AND sort <= ? AND group_parent_id = ? AND cus_user_id = ?", s.X, s.Y, s.F, userId).UpdateColumn("sort", gorm.Expr("sort - ?", 1)).Error
	}
	err = global.GVA_DB.Model(&p).Where("sg_sea_engine_id = ? AND cus_user_id = ?", s.P, userId).Update("sort", s.Y).Error
	return
}

func CopyBookmarkGroup(r request.CopyBookmarkGroupRequest) (list interface{}) {
	var allGroup []model.CusBookmarkGroup
	var g []model.CusBookmarkGroup
	global.GVA_DB.Where("g_sea_engine_id = ?", r.CusGroupID).First(&g)
	g[0].GroupParentID = 0
	global.GVA_DB.Order("sort ASC").Find(&allGroup)
	respNodes := utils.FindRelationNode(model.CusBookmarkGroups.ConvertToINodeArray(g), model.CusBookmarkGroups.ConvertToINodeArray(allGroup))

	//var s []model.CusShareGroup
	for i, _ := range respNodes {
		index := reflect.ValueOf(respNodes).Index(i)
		fmt.Println(index)
	}

	return respNodes
}
