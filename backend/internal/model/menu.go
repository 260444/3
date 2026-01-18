package model

import (
	"gorm.io/gorm"
)

// Menu 菜单模型
type Menu struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"size:50;not null" json:"name"`
	Title     string         `gorm:"size:50;not null" json:"title"`
	Path      string         `gorm:"size:100" json:"path"`
	Component string         `gorm:"size:100" json:"component"`
	Redirect  string         `gorm:"size:100" json:"redirect"`
	ParentID  *uint          `json:"parent_id"` // 父级菜单ID
	Parent    *Menu          `gorm:"foreignKey:ParentID" json:"parent"`
	Children  []Menu         `gorm:"foreignKey:ParentID" json:"children"`
	Icon      string         `gorm:"size:50" json:"icon"`
	Sort      int            `gorm:"default:0" json:"sort"`
	IsHidden  bool           `gorm:"default:false" json:"is_hidden"`
	IsLink    bool           `gorm:"default:false" json:"is_link"`
	LinkUrl   string         `gorm:"size:255" json:"link_url"`
	Status    int            `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	Roles     []Role         `gorm:"many2many:role_menus;" json:"roles"`
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menus"
}