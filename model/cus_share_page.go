package model

import "github.com/siuvlqnm/bookmark/global"

type CusSharePage struct {
	global.GVA_MODEL
	CusUserID    uint   `json:"cusUserId"`
	PSeaEngineID uint32 `json:"pSeaEngineId"`
	PageName     string `json:"pageName"`
	IsPassword   bool   `json:"isPassword"`
	PagePassword string `json:"pagePassword"`
	Serve        string
}
