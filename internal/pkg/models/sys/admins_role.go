package sys

import "github.com/mengjayxc/webmanage-api/internal/pkg/models/basemodel"

// 用户-角色
type AdminsRole struct {
	basemodel.Model
	AdminsID uint64 `gorm:"column:admins_id;unique_index:uk_admins_role_admins_id;not null;"` // 管理员ID
	RoleID   uint64 `gorm:"column:role_id;unique_index:uk_admins_role_admins_id;not null;"`   // 角色ID
}

