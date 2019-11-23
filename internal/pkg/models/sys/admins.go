package sys

import (
	"time"
)

// 后台用户
type Admins struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"`
	UserName  string    `gorm:"column:user_name;size:32;unique_index:uk_admins_user_name;not null;" json:"user_name" form:"user_name"`
	Password  string    `gorm:"column:password;type:char(32);not null;" json:"password" form:"password"`
}






