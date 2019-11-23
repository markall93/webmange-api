package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/controllers/common"
	models "github.com/mengjayxc/webmanage-api/internal/pkg/models/common"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/web"
	"github.com/mengjayxc/webmanage-api/pkg/logger"
	"strconv"
	"time"
)

func AddAccounts(c *gin.Context) {
	accountAddr := c.Query("accountName")
	amt := c.Query("amount")
	amount, _ := strconv.Atoi(amt)
	timeAt := c.Query("bindedAt")
	fmt.Println("timeAt:>>>>>>>>>", timeAt)
	bindedTime, _ := time.Parse(common.TimeLayout, timeAt)

	//bindedTime, _ := time.ParseInLocation(common.TimeLayout, timeAt, time.Local)
	fmt.Println("bindedTime:>>>>>>>>>>>.", bindedTime)

	// 根据当前账户地址，从数据库中查找该账户累计参与次数和累计参与金额
	var infos web.UserInfos
	where := web.UserInfos{}
	where.AccountAddr = accountAddr
	notFind, _ := models.First(&where, &infos)
	if notFind {
		logger.Debug("table: user_infos no data")
		newInfos := web.UserInfos{
			AccountAddr:  accountAddr,
			TotalNumbers: 1,
			TotalAmounts: amount,
		}
		err := models.Create(&newInfos)
		if err != nil {
			common.ResFail(c, "添加userInfos表数据，操作失败")
			return
		}
	} else {
		newInfos := web.UserInfos{
			AccountAddr:  accountAddr,
			TotalNumbers: infos.TotalNumbers + 1,
			TotalAmounts: infos.TotalAmounts + amount,
		}

		err := models.Updates(&infos, &newInfos)
		if err != nil {
			common.ResFail(c, "更新userInfos表数据，操作失败")
			return
		}
	}
	accountModel := web.Accounts{AccountAddr: accountAddr, BindedTime:bindedTime, Amount:amount}
	err := models.Create(&accountModel)
	if err != nil {
		common.ResFail(c, "添加account表数据，操作失败")
		return
	}
	common.ResSuccess(c, gin.H{"msg": "添加数据成功"})
}

// 分页返回参与账户信息
func GetAllAccounts(c *gin.Context) {
	page := common.GetPageIndex(c)
	limit := common.GetPageLimit(c)
	var whereOrder []models.PageWhereOrder
	order := "ID ASC"

	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})

	var total int
	var accountList []web.Accounts
	err := models.GetPage(&web.Accounts{}, &web.Accounts{}, &accountList, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	// 根据当前账户地址，从'UserInfos'表中查找该账户累计参与次数和累计参与金额
	// 去除accountList中重复的元素
	accounts := common.RemoveRepeatedElement(accountList)

	var userInfosList []*web.UserInfos
	for _, account := range accounts {
		var infos web.UserInfos
		where := web.UserInfos{}
		accountAddr := account.AccountAddr
		where.AccountAddr = accountAddr

		_, err := models.First(&where, &infos)
		if err != nil {
			common.ResErrSrv(c, err)
			return
		}
		userInfosList = append(userInfosList, &infos)
	}

	common.ResSuccessPage(c, total, map[string]interface{}{
		"accounts": accountList,
		"totalInfos": userInfosList,
	})
}

//// 用户结算-->数据库
//func Settlement(c *gin.Context) {
//	accountAddr := c.Query("accountName")
//	settle := c.Query("settleTime")
//	settleTime, _ := time.Parse(common.TimeLayout, settle)
//	amount := c.Query("amount")
//	settleAmount, _ := strconv.Atoi(amount)
//
//	// 结算时间 - 7天就是参与时间
//	// 根据地址和参与时间，在账户表中找出该账户的参与金额
//	//func (t Time) Unix() int64
//	//time1 := settleTime.UnixNano()
//
//
//
//
//
//
//
//
//
//
//
//
//	settleModel := web.UserSettle{
//		AccountAddr:  accountAddr,
//		BindedAmount: ,
//		SettleTime:   settleTime,
//		SettleAmount: settleAmount,
//	}
//
//	err := models.Create(&settleModel)
//	if err != nil {
//		common.ResFail(c, "添加user_settle表数据, 操作失败")
//	}
//}

// 数据库-->用户结算信息返回给前端
//func GetSettleinfos(c *gin.Context) {
//	accountAddr := c.Query("accountName")
//	// 根据当前账户地址，从userSettle表中查找该账户的参与记录(结算时间，结算金额)
//
//	// 把参与账户，参与金额， 结算时间， 结算金额
//
//
//
//	//var infos web.UserInfos
//	//where := web.UserInfos{}
//	//where.AccountAddr = accountAddr
//	//notFind, _ := models.First(&where, &infos)
//	//if notFind {
//	//	logger.Debug("table: user_infos no data")
//	//	newInfos := web.UserInfos{
//	//		AccountAddr:  accountAddr,
//	//		TotalNumbers: 1,
//	//		TotalAmounts: amount,
//	//	}
//	//	err := models.Create(&newInfos)
//	//	if err != nil {
//	//		common.ResFail(c, "添加userInfos表数据，操作失败")
//	//		return
//	//	}
//
//
//
//
//
//}




// 累计参与人数：统计'accounts'数据库的总数, totalPartNums
// 累计参与金额：所有账户参与金额的总和, totalPartAmounts
// 累计领取金额：玩家从资金池中领取的总金额 ? , totalWithdrawAmts
func GetTotalInfos(c *gin.Context) {
	var allAccounts []web.Accounts
	//Participate
	var totalPartNums int
	var totalPartAmounts int
	// totalWithdrawAmts

	err := models.GetInfos(&web.Accounts{}, &web.Accounts{}, &allAccounts, &totalPartNums)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}

	for _, account := range allAccounts {
		totalPartAmounts += account.Amount
	}

	common.ResSuccessInfos(c, map[string]interface{}{
		"totalPartNums": totalPartNums,
		"totalPartAmounts": totalPartAmounts,
	})
}


