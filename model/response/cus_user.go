package response

import (
	"github.com/siuvlqnm/bookmark/model"
)

type CusUserResponse struct {
	User model.CusUser `json:"user"`
}

type CusLoginResponse struct {
	User      model.CusUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
