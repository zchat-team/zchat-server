package model

import (
	"time"
)

// 登录日志
type UserLoginLog struct {
	Id        int64     `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"`  // 系统编号
	Uid       int64     `json:"uid" gorm:"column:uid;type:bigint(20);not null;default:0"`       // 用户ID
	Ua        string    `json:"ua" gorm:"column:ua;type:varchar(50);not null"`                  // UserAgent
	DeviceId  string    `json:"device_id" gorm:"column:device_id;type:varchar(64);not null"`    // 设备ID
	Type      int       `json:"type" gorm:"column:type;type:tinyint(4);not null;default:0"`     // 类型[1:登录;2:登出]
	Status    int       `json:"status" gorm:"column:status;type:tinyint(4);not null;default:0"` // 状态[1:成功;2:失败]
	LoginIp   string    `json:"login_ip" gorm:"column:login_ip;type:varchar(50);not null"`      // 登录IP
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`     // 创建时间
}

func (_ *UserLoginLog) TableName() string {
	return "user_login_log"
}
