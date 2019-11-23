package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/middleware"
)

func RegisterRouter(app *gin.Engine) {
	apiv1:="/api/v1"
	apiv2:="/api/v2"
	gv1 := app.Group(apiv1)
	gv2 := app.Group(apiv2)
	// web
	RegisterRouterWeb(gv1)

	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiv2+"/user/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiv2+"/user/logout")
	gv2.Use(middleware.UserAuthMiddleware(
		middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	))

	// sys
	RegisterRouterSys(gv2)
}


