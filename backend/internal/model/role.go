package model

import (
	"gorm.io/gorm"
)

// Role 角色模型
type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"uniqueIndex;size:50;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Status      int            `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	Users       []User         `gorm:"foreignKey:RoleID" json:"users"`
	Menus       []Menu         `gorm:"many2many:role_menus;" json:"menus"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "roles"
}