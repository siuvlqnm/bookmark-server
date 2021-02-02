package request

type SetPageSort struct {
	X int    `json:"x"`
	Y int    `json:"y"`
	F int    `json:"f"`
	P uint32 `json:"p"`
}

type CopyBookmarkGroupRequest struct {
	PSeaEngineID uint32 `json:"pSeaEngineId"`
	CusGroupID   uint32 `json:"groupId"`
	//ShareGroupParentID uint32 `json:"shareGroupParentId"`
}
