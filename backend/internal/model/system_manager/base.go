package system_manager

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
