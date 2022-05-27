package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// 好友申请
type Application struct {
	Id        int64                 `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"`
	Uid       int64                 `json:"uid" gorm:"column:uid;type:bigint(20);not null;default:0"`
	FriendUid int64                 `json:"friend_uid" gorm:"column:friend_uid;type:bigint(20);not null;default:0"`
	Status    int                   `json:"status" gorm:"column:status;type:tinyint(4);not null;default:0"` // 状态[1:审核中 2:通过 3:拒绝]
	IsRead    string                `json:"is_read" gorm:"column:is_read;type:char(1);not null;default:0"`
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`
	ExpiresAt int64                 `json:"expires_at" gorm:"column:expires_at;type:bigint(20);not null;default:0"`
}

func (_ *Application) TableName() string {
	return "application"
}
