package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

type CusInclRecord struct {
	global.GVA_MODEL
	IRSeaEngineID uint32 `json:"irSeaEngineId"`
	InclUserID    uint   `json:"-"`
	SharePageID   uint32 `json:"sharePageId"`
	ShareGroupID  uint32 `json:"groupId"`
	TargetUrl     string `json:"targetUrl"`
	Domain        string `json:"domain"`
	Path          string `json:"path"`
	Query         string `json:"query"`
	Title         string `json:"title"`
	Description   string `json:"description" gorm:"type:varchar(300)"`
	IsAccept      uint8  `json:"isAccept"`
	Reply         string `json:"reply"`
}

func (i *CusInclRecord) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(i).Update("ir_sea_engine_id", utils.GetMurmur32("inclRec:", i.ID)).Error
	return err
}
