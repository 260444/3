package model

import "time"

type CasbinRule struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Ptype     string    `json:"ptype" gorm:"size:10;not null"` //策略类型p(策略)g(角色)
	V0        string    `json:"v0" gorm:"size:100"`            //主体(sub)
	V1        string    `json:"v1" gorm:"size:100"`            //客体(obj)
	V2        string    `json:"v2" gorm:"size:100"`            //动作(act)
	V3        string    `json:"v3" gorm:"size:100"`
	V4        string    `json:"v4" gorm:"size:100"`
	V5        string    `json:"v5" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
