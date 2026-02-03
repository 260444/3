package system_manager

// Menu 菜单模型
type Menu struct {
	BaseModel
	// Name 菜单名称
	Name string `gorm:"size:50;not null" json:"name"`
	// Title 菜单标题
	Title string `gorm:"size:50;not null" json:"title"`
	// Path 路由路径
	Path string `gorm:"size:100" json:"path"`
	// Component 组件路径
	Component string `gorm:"size:100" json:"component"`
	// ParentID 父级菜单ID
	ParentID *uint `json:"parent_id" gorm:"default:0"`
	// Parent 父级菜单
	Parent *Menu `gorm:"-" json:"parent"`
	// Children 子菜单列表
	Children []Menu `gorm:"-" json:"children"`
	// Icon 菜单图标
	Icon string `gorm:"size:50" json:"icon"`
	// Sort 排序值
	Sort int `gorm:"default:0" json:"sort"`
	// IsHidden 是否隐藏
	IsHidden bool `gorm:"default:false" json:"is_hidden"`
	// Status 状态: 1-正常, 0-禁用
	Status int `gorm:"default:1" json:"status"`
}

// TableName 指定表名
func (Menu) TableName() string {
	return "menus"
}
