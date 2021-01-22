package response

import "github.com/siuvlqnm/bookmark/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
