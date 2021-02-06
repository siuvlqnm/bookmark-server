package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

type CusBookmark struct {
	global.GVA_MODEL
	MSeaEngineID uint32 `json:"mSeaEngineId"`
	CusWebID     uint   `json:"-"`
	CusUserID    uint   `json:"-"`
	TargetUrl    string `json:"targetUrl"`
	Domain       string `json:"domain"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	Title        string `json:"title"`
	Description  string `json:"description" gorm:"type:varchar(300)"`
	Icon         string `json:"icon"`
	CusGroupID   uint   `json:"groupId"`
	ShareGroupID uint32 `json:"shareGroupId"`
	CusTagStr    string `json:"cusTagStr"`
	IsStar       uint8  `json:"isStar"`
}

func (b *CusBookmark) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(b).Update("m_sea_engine_id", utils.GetMurmur32("bkm:", b.ID)).Error
	return err
}
