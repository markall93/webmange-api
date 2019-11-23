package basemodel

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
}

