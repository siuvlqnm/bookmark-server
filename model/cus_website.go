package model

import "github.com/siuvlqnm/bookmark/global"

type CusWebsite struct {
	global.GVA_MODEL
	Domain      string `json:"domain"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:varchar(300)"`
}
