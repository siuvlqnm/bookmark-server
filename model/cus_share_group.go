package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
	"gorm.io/gorm"
)

type CusShareGroup struct {
	global.GVA_MODEL
	CusUserID     uint          `json:"-"`
	SharePageID   uint32        `json:"sharePageId"`
	SGSeaEngineID uint32        `json:"sGSeaEngineId"`
	GroupParentID int           `json:"groupParentId"`
	GroupName     string        `json:"groupName"`
	GroupIcon     string        `json:"groupIcon"`
	Sort          int           `json:"-"`
	Bookmark      []CusBookmark `json:"bookmark" gorm:"foreignKey:ShareGroupID;references:SGSeaEngineID"`
}

func (g *CusShareGroup) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(g).Update("sg_sea_engine_id", utils.GetMurmur32("shreGrp:", g.ID)).Error
	return err
}

// region 实现ITree 所有接口
func (g CusShareGroup) GetTitle() string {
	return g.GroupName
}

func (g CusShareGroup) GetId() int {
	return int(g.SGSeaEngineID)
}

func (g CusShareGroup) GetFatherId() int {
	return g.GroupParentID
}

func (g CusShareGroup) GetData() interface{} {
	return g
}

func (g CusShareGroup) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return g.GroupParentID == 0 || uint(g.GroupParentID) == g.ID
}

// endregion

type CusShareGroups []CusShareGroup

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (g CusShareGroups) ConvertToINodeArray() (nodes []utils.INode) {
	for _, v := range g {
		nodes = append(nodes, v)
	}
	return
}
