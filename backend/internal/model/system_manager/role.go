// Package system_manager 定义系统管理相关的数据模型。
//
// 该包包含用户、角色、菜单、权限等核心数据模型的定义。
// 所有模型都使用 GORM 标签进行数据库映射。
package system_manager

// Role 表示系统角色模型。
//
// Role 结构体存储角色的基本信息，包括角色名称、标识符、描述等。
// 角色是 RBAC 权限模型的核心概念，用于将权限分配给用户组。
type Role struct {
	BaseModel
	Name        string `gorm:"uniqueIndex;size:50;not null" json:"name"`  // 角色名称，唯一索引
	Ident       string `gorm:"uniqueIndex;size:50;not null" json:"ident"` // 角色标识符，唯一索引
	Description string `gorm:"size:255" json:"description"`               // 角色描述
	Status      int    `gorm:"default:1" json:"status"`                   // 状态：1-正常，0-禁用
}

// TableName 指定 Role 模型对应的数据库表名。
//
// 返回 "roles" 作为表名，覆盖 GORM 默认的表名生成规则。
func (Role) TableName() string {
	return "roles"
}
