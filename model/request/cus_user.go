package request

import uuid "github.com/satori/go.uuid"

// User register structure
type CRegister struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	NickName string `json:"nickName"`
}

// User login structure
type CLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Modify password structure
type CChangePasswordStruct struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type CSetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
