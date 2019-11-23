package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/api/v1"
	"github.com/mengjayxc/webmanage-api/api/v2"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//apiPrefix := "/api"
	//g := r.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	//var notCheckLoginUrlArr []string




	r.POST("/user/register", v2.Register)
	r.POST("/user/login", v2.Login)
	r.GET("/user/logout", v2.Logout)
	r.POST("/user/editpwd", v2.EditPwd)
	r.POST("/addAccount", v1.AddAccount)


	

	return r
}


