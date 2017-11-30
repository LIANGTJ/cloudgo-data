package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int `xorm:"pk autoincr"` //语义标签
	Username   string
	Departname string
	CreatedAt  time.Time `xorm:"created"`
}

// func (UserInfo*) TableName() string {
// 	return "UserInfo"
// }

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.Username) == 0 {
		panic("Username shold not null!")
	}
	return &u
}
