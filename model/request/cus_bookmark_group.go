package request

type PostBookmarkGroup struct {
	GSeaEngineId uint32 `json:"gSeaEngineId"`
	IsArchive    bool   `json:"isArchive"`
}

type GetGetBookmarkGroup struct {
	GSeaEngineId uint32 `uri:"gSeaEngineId"`
	IsArchive    bool   `form:"isArchive"`
}

type SetGroupSort struct {
	X int    `json:"x"`
	Y int    `json:"y"`
	F int    `json:"f"`
	G uint32 `json:"g"`
}
