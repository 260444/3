package system_manager

// Role 角色模型
type Role struct {
	BaseModel
	Name        string `gorm:"uniqueIndex;size:50;not null" json:"name"`
	Ident       string `gorm:"uniqueIndex;size:50;not null" json:"ident"`
	Description string `gorm:"size:255" json:"description"`
	Status      int    `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
}

// TableName 指定表名
func (Role) TableName() string {
	return "roles"
}
