package model

// OperationLog 操作日志模型
type OperationLog struct {
	BaseModel
	Operation     string `gorm:"size:255;not null" json:"operation"`
	UserID        uint   `gorm:"not null" json:"user_id"`
	Username      string `gorm:"size:50;not null" json:"username"`
	IP            string `gorm:"size:45" json:"ip"`
	UserAgent     string `gorm:"size:500" json:"user_agent"`
	Status        int    `gorm:"default:1" json:"status"` // 1: 成功, 0: 失败
	RequestMethod string `gorm:"size:10" json:"request_method"`
	RequestPath   string `gorm:"size:255" json:"request_path"`
	RequestBody   string `gorm:"type:text" json:"request_body"`
	ResponseBody  string `gorm:"type:text" json:"response_body"`
	ResponseTime  int64  `json:"response_time"` // 响应时间(毫秒)
}

// TableName 指定表名
func (OperationLog) TableName() string {
	return "operation_logs"
}
