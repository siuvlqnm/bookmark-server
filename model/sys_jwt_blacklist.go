package model

import (
	"github.com/siuvlqnm/bookmark/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
