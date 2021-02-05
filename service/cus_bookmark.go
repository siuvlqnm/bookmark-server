package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
)

func CreateBookmark(b model.CusBookmark) (err error, bookmark *model.CusBookmark) {
	err = global.GVA_DB.Create(&b).Error
	return err, &b
}

func GetBookmarkList(userId uint, where model.CusBookmark, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CusBookmark{})
	var bookmarkList []model.CusBookmark

	db = db.Where("cus_user_id = ?", userId)

	if where.CusGroupID != 0 {
		db = db.Where("cus_group_id = ?", where.CusGroupID)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, bookmarkList, total
	} else {
		err = db.Order("id desc").Limit(limit).Offset(offset).Find(&bookmarkList).Error
	}
	return err, bookmarkList, total
}

func UpdateBookmarkMSeaEngineId(id int, MSeaEngineId uint32) {
	var bookmark model.CusBookmark
	global.GVA_DB.Model(&bookmark).Where("id = ?", id).Update("m_sea_engine_id", MSeaEngineId)
	return
}

func UpdateBookmark(MSeaEngineId uint32, b *model.CusBookmark) (err error) {
	var bookmark *model.CusBookmark
	upDateMap := make(map[string]interface{})
	upDateMap["cus_web_id"] = b.CusWebID
	upDateMap["domain"] = b.Domain
	upDateMap["path"] = b.Path
	upDateMap["query"] = b.Query
	upDateMap["title"] = b.Title
	upDateMap["description"] = b.Description
	upDateMap["cus_tag_str"] = b.CusTagStr
	upDateMap["cus_group_id"] = b.CusGroupID
	upDateMap["is_star"] = b.IsStar
	upDateMap["share_group_id"] = b.ShareGroupID
	err = global.GVA_DB.Model(&bookmark).Where("m_sea_engine_id = ?", MSeaEngineId).Updates(upDateMap).Error
	return
}

func DeleteBookmark(MSeaEngineId uint32) (err error) {
	var b model.CusBookmark
	err = global.GVA_DB.Where("m_sea_engine_id = ?", MSeaEngineId).Delete(&b).Error
	return
}

func UpdateToStar(MSeaEngineId uint32, IsStar uint8) (err error) {
	var bookmark model.CusBookmark
	err = global.GVA_DB.Model(&bookmark).Where("m_sea_engine_id = ?", MSeaEngineId).Update("is_star", IsStar).Error
	return
}
