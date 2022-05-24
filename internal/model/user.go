package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// 用户
type User struct {
	Id             int64                 `json:"id" gorm:"primaryKey;column:id;type:bigint(20)"`                                   // 用户ID
	Zid            string                `json:"zid" gorm:"column:zid;type:varchar(50);not null"`                                  // 知聊号
	Password       string                `json:"password" gorm:"column:password;type:varchar(64);not null"`                        // 密码
	CreatedAt      time.Time             `json:"created_at" gorm:"column:created_at;type:datetime;not null"`                       // 创建时间
	UpdatedAt      time.Time             `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`                       // 更新时间
	DeletedAt      soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`           // 删除时间
	Status         int                   `json:"status" gorm:"column:status;type:tinyint(4);not null;default:0"`                   // 状态
	Nickname       string                `json:"nickname" gorm:"column:nickname;type:varchar(50);not null"`                        // 昵称
	Mobile         string                `json:"mobile" gorm:"column:mobile;type:varchar(50);not null"`                            // 手机号
	MobileVerified int                   `json:"mobile_verified" gorm:"column:mobile_verified;type:tinyint(4);not null;default:0"` // 手机是否被验证
	Email          string                `json:"email" gorm:"column:email;type:varchar(50);not null"`                              // 邮箱
	EmailVerified  int                   `json:"email_verified" gorm:"column:email_verified;type:tinyint(4);not null;default:0"`   // 邮箱是否被验证
	Avatar         string                `json:"avatar" gorm:"column:avatar;type:varchar(256);not null"`                           // 头像
	Gender         int                   `json:"gender" gorm:"column:gender;type:tinyint(4);not null;default:0"`                   // 性别[1:男;2:女]
	RegisterIp     string                `json:"register_ip" gorm:"column:register_ip;type:varchar(50);not null"`                  // 注册IP
	LastLoginTime  time.Time             `json:"last_login_time" gorm:"column:last_login_time;type:datetime"`                      // 最近登录时间
	LastLoginIp    string                `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(50);not null"`              // 最近登录IP
	LoginIpLimit   string                `json:"login_ip_limit" gorm:"column:login_ip_limit;type:text;not null"`                   // 登录限制,JSON格式
}

func (_ *User) TableName() string {
	return "user"
}
