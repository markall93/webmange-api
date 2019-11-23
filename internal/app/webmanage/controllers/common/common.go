package common

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/web"
	"net/http"
)

const (
	SuccessCode         = 2000       //成功的状态码
	FailCode            = 4000       //失败的状态码
	Md5Prefix           = "jkfldfsf" //MD5加密前缀字符串
	TokenKey            = "X-Token"  //页面token键名
	UserIdKey           = "X-USERID" //页面用户ID键名
	UserUuidKey         = "X-UUID"   //页面UUID键名
	SuperAdminId uint64 = 956896     // 超级管理员账号ID
    TimeLayout          = "2006-01-02 15:04:05"
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SuccessCode, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModelBase{Code: SuccessCode, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: FailCode, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}

// 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FailCode, Message: "服务端故障"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FailCode, Message: "err"}
	ResJSON(c, http.StatusOK, &ret)
}

type ResponsePageData struct {
	Total int      `json:"total"`
	Items interface{} `json:"items"`
}



type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}
type ResponseInfoData struct {
	Items interface{} `json:"items"`
}

type ResponseInfos struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponseInfoData `json:"data"`
}

// 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total int, list interface{}) {
	ret := ResponsePage{Code: SuccessCode, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功-totalInfos
func ResSuccessInfos(c *gin.Context, data interface{}) {
	ret := ResponseInfos{Code: SuccessCode, Message: "ok", Data: ResponseInfoData{Items: data}}
	ResJSON(c, http.StatusOK, &ret)
}


// 获取页码
func GetPageIndex(c *gin.Context) uint64 {
	return GetQueryToUint64(c, "page", 1)
}

// 获取每页记录数
func GetPageLimit(c *gin.Context) uint64 {
	limit := GetQueryToUint64(c, "limit", 20)
	if limit > 500 {
		limit = 20
	}
	return limit
}

// 检查用户是否有权限
func CsbinCheckPermission(userID, url, methodtype string) (bool, error) {
	return Enforcer.EnforceSafe(PrefixUserID+userID, url, methodtype)
}

// 去除切片中重复元素
func RemoveRepeatedElement(arr []web.Accounts) (newArr []web.Accounts) {
	newArr = make([]web.Accounts, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i].AccountAddr == arr[j].AccountAddr {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}


