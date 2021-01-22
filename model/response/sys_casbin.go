package response

import "github.com/siuvlqnm/bookmark/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
