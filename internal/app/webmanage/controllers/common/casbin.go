package common

import (
	"github.com/casbin/casbin"
	models "github.com/mengjayxc/webmanage-api/internal/pkg/models/common"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/sys"
	"github.com/mengjayxc/webmanage-api/pkg/convert"
)

const (
	PrefixUserID = "u"
	PrefixRoleID = "r"
)

var Enforcer *casbin.Enforcer

// 用户角色处理
func CsbinAddRoleForUser(userid uint64)(err error){
	if Enforcer == nil {
		return
	}
	uid:=PrefixUserID+convert.ToString(userid)
	Enforcer.DeleteRolesForUser(uid)
	var adminsroles []sys.AdminsRole
	err = models.Find(&sys.AdminsRole{AdminsID: userid}, &adminsroles)
	if err != nil {
		return
	}
	for _, adminsrole := range adminsroles {
		Enforcer.AddRoleForUser(uid, PrefixRoleID+convert.ToString(adminsrole.RoleID))
	}
	return
}
