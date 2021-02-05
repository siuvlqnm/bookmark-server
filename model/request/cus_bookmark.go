package request

import "github.com/siuvlqnm/bookmark/model"

type NewBookmark struct {
	MSeaEngineId uint32 `json:"mSeaEngineId"`
	Link         string `json:"link"`
	Domain       string `json:"domain"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TagStr       string `json:"tagStr"`
	CusGroupId   uint32 `json:"groupId"`
	IsStar       uint8  `json:"isStar"`
}

type GetBookmarkList struct {
	PageInfo
	model.CusBookmark
}

type CopyBookmark struct {
	MSeaEngineId uint32 `json:"mSeaEngineId"`
	ShareGroupID uint32 `json:"shareGroupId"`
}
