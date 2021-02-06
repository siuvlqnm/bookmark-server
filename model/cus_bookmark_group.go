package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

type CusBookmarkGroup struct {
	global.GVA_MODEL
	CusUserID     uint          `json:"-"`
	GSeaEngineID  uint32        `json:"gSeaEngineId"`
	GroupParentID int           `json:"groupParentId"`
	GroupName     string        `json:"groupName"`
	GroupIcon     string        `json:"groupIcon"`
	Sort          int           `json:"-"`
	IsArchive     bool          `json:"isArchive"`
	Bookmark      []CusBookmark `json:"bookmark" gorm:"foreignKey:CusGroupID;references:GSeaEngineID"`
}

func (g *CusBookmarkGroup) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(g).Update("g_sea_engine_id", utils.GetMurmur32("bkmGrp:", g.ID)).Error
	return err
}

// region 实现ITree 所有接口
func (g CusBookmarkGroup) GetTitle() string {
	return g.GroupName
}

func (g CusBookmarkGroup) GetId() int {
	return int(g.GSeaEngineID)
}

func (g CusBookmarkGroup) GetFatherId() int {
	return g.GroupParentID
}

func (g CusBookmarkGroup) GetData() interface{} {
	return g
}

func (g CusBookmarkGroup) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return g.GroupParentID == 0 || uint(g.GroupParentID) == g.ID
}

// endregion

type CusBookmarkGroups []CusBookmarkGroup

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (g CusBookmarkGroups) ConvertToINodeArray() (nodes []utils.INode) {
	for _, v := range g {
		nodes = append(nodes, v)
	}
	return
}
