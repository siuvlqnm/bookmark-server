package request

type PostBookmarkGroup struct {
	GSeaEngineId uint32 `json:"gSeaEngineId"`
	GShareId     uint32 `json:"gShareId"`
	IsShare      bool   `json:"isShare"`
	IsArchive    bool   `json:"isArchive"`
}

type GetGetBookmarkGroup struct {
	IsArchive bool `form:"isArchive"`
}
