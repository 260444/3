// Package system_manager 定义系统管理相关的数据模型。
//
// 该包包含用户、角色、菜单、权限等核心数据模型的定义。
// 所有模型都使用 GORM 标签进行数据库映射。
package system_manager

// Menu 表示系统菜单模型。
//
// Menu 结构体存储菜单的基本信息，支持树形结构。
// 菜单用于构建系统的导航结构，可以关联到角色以控制访问权限。
type Menu struct {
	BaseModel
	Name      string `gorm:"size:50;not null" json:"name"`   // 菜单名称
	Title     string `gorm:"size:50;not null" json:"title"`  // 菜单标题
	Path      string `gorm:"size:100" json:"path"`           // 路由路径
	Component string `gorm:"size:100" json:"component"`      // 组件路径
	ParentID  *uint  `json:"parent_id" gorm:"default:0"`     // 父级菜单 ID
	Parent    *Menu  `gorm:"-" json:"parent"`                // 父级菜单（不映射到数据库）
	Children  []Menu `gorm:"-" json:"children"`              // 子菜单列表（不映射到数据库）
	Icon      string `gorm:"size:50" json:"icon"`            // 菜单图标
	Sort      int    `gorm:"default:0" json:"sort"`          // 排序值
	IsHidden  bool   `gorm:"default:false" json:"is_hidden"` // 是否隐藏
	Status    int    `gorm:"default:1" json:"status"`        // 状态：1-正常，0-禁用
}

// TableName 指定 Menu 模型对应的数据库表名。
//
// 返回 "menus" 作为表名，覆盖 GORM 默认的表名生成规则。
func (Menu) TableName() string {
	return "menus"
}
