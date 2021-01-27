package response

type CusBookmarkGroupResponse struct {
	CusUserId     uint                           `json:"-"`
	GSeaEngineId  uint32                         `json:"gSeaEngineId"`
	GroupParentId int                            `json:"groupParentId"`
	GroupName     string                         `json:"groupName"`
	GroupIcon     string                         `json:"groupIcon"`
	IsArchive     bool                           `json:"isArchive"`
	Children      []CusGroupWithBookmarkResponse `json:"children" gorm:"-"`
}

type CusGroupWithBookmarkResponse struct {
	CusUserId     uint                           `json:"-"`
	GSeaEngineId  uint32                         `json:"gSeaEngineId"`
	GroupParentId int                            `json:"groupParentId"`
	GroupName     string                         `json:"groupName"`
	GroupIcon     string                         `json:"groupIcon"`
	IsArchive     bool                           `json:"isArchive"`
	Children      []CusGroupWithBookmarkResponse `json:"children" gorm:"-"`
}
