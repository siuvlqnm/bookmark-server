package model

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/utils"
)

type CusBookmarkGroup struct {
	global.GVA_MODEL
	CusUserId     uint               `json:"-"`
	GSeaEngineId  uint32             `json:"gSeaEngineId"`
	GroupParentId int                `json:"-"`
	GroupName     string             `json:"groupName"`
	GroupIcon     string             `json:"groupIcon"`
	Sort          int                `json:"-"`
	IsArchive     bool               `json:"isArchive"`
	Children      []CusBookmarkGroup `json:"children" gorm:"-"`
	Bookmark      []CusBookmark      `json:"bookmark" gorm:"foreignKey:CusGroupId"`
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
	return g.GroupParentId == 0 || g.GroupParentId == int(g.ID)
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
