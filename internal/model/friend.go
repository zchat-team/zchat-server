package model

import (
	"time"
)

// 朋友
type Friend struct {
	Id        int64     `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"`
	Uid       int64     `json:"uid" gorm:"column:uid;type:bigint(20);not null"`
	FriendUid int64     `json:"friend_uid" gorm:"column:friend_uid;type:bigint(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
	Alias     string    `json:"alias" gorm:"column:alias;type:varchar(50);not null"`
}

func (_ *Friend) TableName() string {
	return "friend"
}
