package model

import (
	"time"
)

type BaseModel struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	Deleted   int       `gorm:"column:deleted;default:0" json:"deleted"`
}
