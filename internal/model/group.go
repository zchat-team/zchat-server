package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Group struct {
	Id           int64                 `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"` // 群ID
	Owner        string                `json:"owner" gorm:"column:owner;type:varchar(50);not null"`           // 群主
	Type         int                   `json:"type" gorm:"column:type;type:tinyint(4);not null;default:0"`    // 群类型
	Name         string                `json:"name" gorm:"column:name;type:varchar(50);not null"`             // 群名称
	CreatedAt    time.Time             `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt    time.Time             `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`
	Notice       string                `json:"notice" gorm:"column:notice;type:varchar(200);not null"`
	Introduction string                `json:"introduction" gorm:"column:introduction;type:varchar(200);not null"`
	Avatar       string                `json:"avatar" gorm:"column:avatar;type:varchar(200);not null"`
}

func (_ *Group) TableName() string {
	return "group"
}
