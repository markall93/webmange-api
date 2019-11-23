package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/controllers/common"
	"github.com/mengjayxc/webmanage-api/pkg/cache"
	"github.com/mengjayxc/webmanage-api/pkg/convert"
	"github.com/mengjayxc/webmanage-api/pkg/jwt"
	"strconv"
	"time"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		var uuid string
		if t := c.GetHeader(common.TokenKey); t != "" {
			userInfo,ok:=jwt.ParseToken(t)
			if !ok {
				common.ResFailCode(c,"token 无效",5008)
				return
			}
			exptimestamp, _ := strconv.ParseInt(userInfo["exp"], 10, 64)
			exp := time.Unix(exptimestamp, 0)
			ok=exp.After(time.Now())
			if !ok {
				common.ResFailCode(c,"token 过期",5014)
				return
			}
			uuid = userInfo["uuid"]
		}

		if uuid != "" {
			//查询用户ID
			val,err:=cache.Get([]byte(uuid))
			if err!=nil {
				common.ResFailCode(c,"token 无效",5008)
				return
			}
			userID:=convert.ToUint(string(val))
			c.Set(common.UserUuidKey, uuid)
			c.Set(common.UserIdKey, userID)
		}
		if uuid == "" {
			common.ResFailCode(c,"用户未登录",5008)
			return
		}
	}
}

