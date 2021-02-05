package model

import "github.com/siuvlqnm/bookmark/global"

type CusMaintRecord struct {
	global.GVA_MODEL
	MaintID      uint   `json:"maintId"`
	SharePageID  uint32 `json:"sharePageId"`
	ShareGroupID uint32 `json:"groupId"`
	TargetUrl    string `json:"targetUrl"`
	Domain       string `json:"domain"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	Title        string `json:"title"`
	Description  string `json:"description" gorm:"type:varchar(300)"`
	IsAccept     bool   `json:"isAccept"`
	Reply        string `json:"reply"`
}
