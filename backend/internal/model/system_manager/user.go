// Package system_manager 定义系统管理相关的数据模型。
//
// 该包包含用户、角色、菜单、权限等核心数据模型的定义。
// 所有模型都使用 GORM 标签进行数据库映射。
package system_manager

import (
	"time"
)

// User 表示系统用户模型。
//
// User 结构体存储用户的基本信息，包括用户名、密码、邮箱、手机号等。
// 密码使用 bcrypt 加密存储。
// 支持用户状态管理（正常/禁用）和角色关联。
type User struct {
	BaseModel
	Username    string     `gorm:"uniqueIndex;size:50;not null" json:"username"` // 用户名，唯一索引
	Password    string     `gorm:"size:255;not null" json:"password"`            // 密码（bcrypt 加密）
	Email       string     `gorm:"uniqueIndex;size:100" json:"email"`            // 邮箱，唯一索引
	Phone       string     `gorm:"size:20" json:"phone"`                         // 手机号
	Nickname    string     `gorm:"size:50" json:"nickname"`                      // 昵称
	Avatar      string     `gorm:"size:255" json:"avatar"`                       // 头像 URL
	Status      int        `gorm:"default:1" json:"status"`                      // 状态：1-正常，0-禁用
	LastLoginAt *time.Time `json:"last_login_at"`                                // 最后登录时间
	LastLoginIP string     `gorm:"size:45" json:"last_login_ip"`                 // 最后登录 IP
	RoleID      *uint      `json:"role_id"`                                      // 关联的角色 ID
}

// UserWithRoleInfo 包含用户和角色信息的结构体。
//
// 该结构体在 User 的基础上扩展了 RoleIdent 字段，用于在查询时同时获取用户的角色标识符。
// 通常用于需要展示用户角色信息的场景。
type UserWithRoleInfo struct {
	User
	RoleIdent string `json:"ident" gorm:"column:ident"` // 角色标识符
}

// TableName 指定 User 模型对应的数据库表名。
//
// 返回 "users" 作为表名，覆盖 GORM 默认的表名生成规则。
func (User) TableName() string {
	return "users"
}
