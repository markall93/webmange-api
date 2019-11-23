package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/api/v1"
)

func RegisterRouterWeb(app *gin.RouterGroup) {
	// 获取网页参与账户信息，插入到数据库中
	app.POST("/account/addaccount", v1.AddAccounts)

	//// 用户结算--> 数据库
	//app.POST("/account/settlement", v1.Settlement)
	//
	//// 数据库--> 用户结算信息返回给前端
	//app.GET("/account/settleinfos", v1.GetSettleinfos)

	// 获取所有的参与账户信息，要分页返回
	app.GET("/account/allaccounts", v1.GetAllAccounts)

	// 综合信息
	// 累计参与人数
	// 累计参与金额
	// 累计提取金额?
	app.GET("/totalinfos", v1.GetTotalInfos)




}

