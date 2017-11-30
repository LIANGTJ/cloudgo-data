package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int `xorm:"pk autoincr"` //语义标签
	UserName   string
	DepartName string
	CreatedAt  time.Time `xorm:"created"`
}



// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("Username shold not null!")
	}
	return &u
}
