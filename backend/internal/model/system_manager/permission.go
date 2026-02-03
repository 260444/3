package system_manager

type Permission struct {
	BaseModel
	Path        string `json:"path" gorm:"type:varchar(255);not null;comment:请求路径"`
	Method      string `json:"method" gorm:"type:varchar(10);not null;comment:请求方法"`
	Description string `json:"description" gorm:"type:varchar(255);comment:权限描述"`
	Status      int8   `json:"status" gorm:"type:tinyint;default:1;comment:请求路径"`
}

func (Permission) TableName() string {
	return "permission"
}
