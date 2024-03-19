package model

type User struct {
	UID           int64  `json:"uid" gorm:"unique"`
	Name          string `json:"name"`
	Sex           int    `json:"sex"`
	Avatar        string `json:"avatar"`
	FollowerCount int64  `json:"fans_count,omitempty"`
	*Medal        `json:"medal,omitempty" gorm:"embedded;embeddedPrefix:medal_"`
}
