// Package system_manager 定义系统管理相关的数据模型。
//
// 该包包含用户、角色、菜单、权限等核心数据模型的定义。
// 所有模型都使用 GORM 标签进行数据库映射。
package system_manager

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 是所有模型的基类，提供通用字段。
//
// 该结构体包含 ID、创建时间、更新时间和删除时间等通用字段。
// 所有其他模型都应该嵌入该结构体以继承这些字段。
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`             // 主键 ID
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`          // 软删除时间
}
