package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Email     string         `gorm:"uniqueIndex;size:100" json:"email"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Nickname  string         `gorm:"size:50" json:"nickname"`
	Avatar    string         `gorm:"size:255" json:"avatar"`
	Status    int            `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	LastLoginAt *time.Time   `json:"last_login_at"`
	LastLoginIP string       `gorm:"size:45" json:"last_login_ip"`
	RoleID    uint           `json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleID" json:"role"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}