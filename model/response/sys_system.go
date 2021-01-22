package response

import "github.com/siuvlqnm/bookmark/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
