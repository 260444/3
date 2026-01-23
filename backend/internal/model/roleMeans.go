package model // RoleMenu 角色菜单关联表
type RoleMenu struct {
	RoleID uint `gorm:"primaryKey" json:"role_id"`
	MenuID uint `gorm:"primaryKey" json:"menu_id"`
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "role_menus"
}

type RoleMenuRequest struct {
	RoleID uint  `json:"role_id"`
	MenuID *uint `json:"menu_id"`
}
