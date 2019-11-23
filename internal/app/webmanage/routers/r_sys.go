package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/api/v2"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	app.POST("/user/register", v2.Register)
	app.POST("/user/login", v2.Login)
	app.POST("/user/logout", v2.Logout)
	app.POST("/user/editpwd", v2.EditPwd)
	//app.POST("/admin/withdraw", v2.WithDrawAsset)
}

