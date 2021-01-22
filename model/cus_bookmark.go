package model

import (
	"github.com/siuvlqnm/bookmark/global"
)

type CusBookmark struct {
	global.GVA_MODEL
	MSeaEngineId uint32 `json:"mSeaEngineId"`
	CusWebId     uint   `json:"-"`
	CusUserId    uint   `json:"-"`
	Protocol     string `json:"protocol"`
	Domain       string `json:"domain"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	CusGroupId   uint   `json:"-"`
	CusTagStr    string `json:"cusTagStr"`
	IsStar       uint8  `json:"isStar"`
}
