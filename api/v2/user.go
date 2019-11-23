package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/controllers/common"
	models "github.com/mengjayxc/webmanage-api/internal/pkg/models/common"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/sys"
	"github.com/mengjayxc/webmanage-api/pkg/cache"
	"github.com/mengjayxc/webmanage-api/pkg/convert"
	"github.com/mengjayxc/webmanage-api/pkg/hash"
	"github.com/mengjayxc/webmanage-api/pkg/jwt"
	"github.com/mengjayxc/webmanage-api/pkg/logger"
	"github.com/mengjayxc/webmanage-api/pkg/util"
	"time"
)

func Register(c *gin.Context) {
	name := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if name == "" || passwd == "" {
		common.ResFail(c, "用户名或密码不能为空")
		return
	}
	insert := sys.Admins{UserName:name, Password:passwd}
	insert.Password = hash.Md5String(common.Md5Prefix + insert.Password)
	err := models.Create(&insert)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccess(c, gin.H{"id": insert.ID})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	
	if username == "" || password == "" {
		common.ResFail(c, "用户名或密码不能为空")
		return
	}
	password = hash.Md5String(common.Md5Prefix + password)
	where := sys.Admins{UserName: username, Password: password}
	user := sys.Admins{}

	// 初始：username = "etherwin, password = 123456
	if username == "etherwin" && password == "900963658df8cd586cf9f31fe665acf7" {
		user.ID = common.SuperAdminId
	} else {
		notFound, err := models.First(&where, &user)
		if err != nil {
			if notFound {
				common.ResFail(c, "用户名或密码错误")
				return
			}
			common.ResErrSrv(c, err)
			return
		}
	}
	//if user.Status != 1 {
	//	common.ResFail(c, "该用户已被禁用")
	//	return
	//}

	// 缓存或者redis
	uuid := util.GetUUID()
	err := cache.Set([]byte(uuid), []byte(convert.ToString(user.ID)), 60*60) // 1H
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	
	// token jwt
	userInfo := make(map[string]string)
	userInfo["exp"] = convert.ToString(time.Now().Add(time.Hour * time.Duration(1)).Unix()) // 1H
	userInfo["iat"] = convert.ToString(time.Now().Unix())
	userInfo["uuid"] = uuid
	token := jwt.CreateToken(userInfo)

	// 发至页面
	resData := make(map[string]string)
	resData["token"] = token
	//casbin 处理
	//err = common.CsbinAddRoleForUser(user.ID)
	//if err != nil {
	//	common.ResErrSrv(c, err)
	//	return
	//}
	logger.Debug("validate login successfully" + "  " + "userID:" + convert.ToString(user.ID))
	common.ResSuccess(c, &resData)
}

// 用户登出, 该怎么处理？
func Logout(c *gin.Context) {
	// 删除缓存
	uuid, exists := c.Get(common.UserUuidKey)
	fmt.Println("uuid>>>>>>>>:", uuid)
	fmt.Println("exists>>>>>>>>:", exists)
	//uuid := "0d63b1c3-5ef5-45bf-aa9d-50e75bd46f3e"

	if exists {
		cache.Del([]byte(convert.ToString(uuid)))
	}
	logger.Debug("logout","UUID:" + "  " + convert.ToString(uuid))
	common.ResSuccessMsg(c)
}

// 用户修改密码
func EditPwd(c *gin.Context) {
	// 用户ID, 怎么获取？
	uid, isExit := c.Get(common.UserIdKey)
	fmt.Println("uid+++++++++:", uid)
	fmt.Println("isExit+++++++++:", isExit)
	if !isExit {
		common.ResFailCode(c, "token 无效", 50008)
		return
	}

	//var uid = 956896
	userID := convert.ToUint64(uid)

	oldPassword := c.Query("old_password")
	newPassword := c.Query("new_password")
	oldPassword = hash.Md5String(common.Md5Prefix + oldPassword)

	if len(newPassword)<6 || len(newPassword)>20 {
		common.ResFail(c, "密码长度在 6 到 20 个字符")
		return
	}
	newPassword = hash.Md5String(common.Md5Prefix + newPassword)
	where := sys.Admins{}
	where.ID = userID
	modelOld := sys.Admins{}
	_, err := models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	if oldPassword != modelOld.Password {
		common.ResFail(c, "原密码输入不正确")
		return
	}
	modelNew:=sys.Admins{Password: newPassword}
	err = models.Updates(&modelOld, &modelNew)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	logger.Debug("editPwd successfully" + "  " + "userID:" + convert.ToString(uid))

	common.ResSuccessMsg(c)
}

//func WithDrawAsset(c *gin.Context) {
//	contractAddr := c.Query("contractAddr")
//
//	adminAddr := c.Query("toAddress")
//
//
//
//
//
//}



