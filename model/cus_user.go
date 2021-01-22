package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/siuvlqnm/bookmark/global"
)

type CusUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"userName"`
	NickName string    `json:nickName`
	Password string    `json:"-"`
	IsVip    uint      `json:"isVip"`
}
