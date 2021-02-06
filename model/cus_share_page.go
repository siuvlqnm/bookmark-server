package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

type CusSharePage struct {
	global.GVA_MODEL
	CusUserID    uint   `json:"-"`
	PSeaEngineID uint32 `json:"pSeaEngineId"`
	PageName     string `json:"pageName"`
	IsPassword   bool   `json:"isPassword"`
	PagePassword string `json:"-"`
	Sort         int    `json:"-"`
}

func (p *CusSharePage) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(p).Update("p_sea_engine_id", utils.GetMurmur32("shrePge:", p.ID)).Error
	return err
}
