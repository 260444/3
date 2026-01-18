package model

import (
	"gorm.io/gorm"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RoleMenu 角色菜单关联表
type RoleMenu struct {
	RoleID uint `gorm:"primaryKey" json:"role_id"`
	MenuID uint `gorm:"primaryKey" json:"menu_id"`
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "role_menus"
}