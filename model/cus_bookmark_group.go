package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
)

type CusBookmarkGroup struct {
	global.GVA_MODEL
	CusUserId     uint               `json:"-"`
	GSeaEngineId  uint32             `json:"gSeaEngineId"`
	GShareId      uint32             `json:"gShareId"`
	GroupParentId int                `json:"groupParentId"`
	GroupName     string             `json:"groupName"`
	GroupIcon     string             `json:"groupIcon"`
	IsArchive     bool               `json:"isArchive"`
	IsShare       bool               `json:"isShare"`
	Children      []CusBookmarkGroup `json:"children" gorm:"-"`
}

// region 实现ITree 所有接口
func (g CusBookmarkGroup) GetTitle() string {
	return g.GroupName
}

func (g CusBookmarkGroup) GetId() int {
	return int(g.ID)
}

func (g CusBookmarkGroup) GetFatherId() int {
	return g.GroupParentId
}

func (g CusBookmarkGroup) GetData() interface{} {
	return g
}

func (g CusBookmarkGroup) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return g.GroupParentId == 0
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
