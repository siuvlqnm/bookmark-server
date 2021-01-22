package model

import "github.com/siuvlqnm/bookmark/global"

type CusBookmarkGroup struct {
	global.GVA_MODEL
	CusUserId     uint   `json:"-"`
	GSeaEngineId  uint32 `json:"gSeaEngineId"`
	GShareId      uint32 `json:"gShareId"`
	GroupParentId int    `json:"-"`
	GroupName     string `json:"groupName"`
	GroupIcon     string `json:"groupIcon"`
	IsArchive     bool   `json:"isArchive"`
	IsShare       bool   `json:isShare`
}
