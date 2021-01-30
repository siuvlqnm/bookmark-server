package model

import "github.com/siuvlqnm/bookmark/global"

type CusShareGroup struct {
	global.GVA_MODEL
	CusUserID     uint          `json:"-"`
	SharePageID   uint32        `json:"sharePageId"`
	SGSeaEngineID uint32        `json:"sGSeaEngineId"`
	GroupParentID int           `json:"groupParentId"`
	GroupName     string        `json:"groupName"`
	GroupIcon     string        `json:"groupIcon"`
	Sort          int           `json:"-"`
	Bookmark      []CusBookmark `json:"bookmark" gorm:"foreignKey:CusGroupId;references:SGSeaEngineId"`
}
