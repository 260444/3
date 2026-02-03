package system_manager

import (
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Username    string     `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password    string     `gorm:"size:255;not null" json:"password"`
	Email       string     `gorm:"uniqueIndex;size:100" json:"email"`
	Phone       string     `gorm:"size:20" json:"phone"`
	Nickname    string     `gorm:"size:50" json:"nickname"`
	Avatar      string     `gorm:"size:255" json:"avatar"`
	Status      int        `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	LastLoginAt *time.Time `json:"last_login_at"`
	LastLoginIP string     `gorm:"size:45" json:"last_login_ip"`
	RoleID      *uint      `json:"role_id"`
}

// UserWithRoleInfo 包含用户和角色信息的结构
type UserWithRoleInfo struct {
	User
	RoleIdent string `json:"ident" gorm:"column:ident"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
