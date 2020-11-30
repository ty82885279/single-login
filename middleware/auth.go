package middleware

import (
	"fmt"
	"single-login/controller"
	"single-login/dao/redis"
	"single-login/pkg/jwt"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		userID := c.PostForm("userID")
		fmt.Printf("UID1111---%v\n", userID)
		if authHeader == "" {

			controller.ResponseErrorWithMsg(c, "请求中缺少Auth Token")
			c.Abort()
			zap.L().Error("请求中缺少Auth Token")
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			controller.ResponseErrorWithMsg(c, "请求头中Token格式错误")
			zap.L().Error("请求头中Token格式错误")
			c.Abort()
			return
		}

		// 从redis取出atoken做对比，
		// 如果不一致,返回错误
		// 如果一致，放行
		//
		AToken, err := redis.GetAToken(userID)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 10086,
				"msg":  "Atoken过期",
			})
			c.Abort()
			return
		}
		fmt.Printf("---%v\n", AToken)
		if AToken != parts[1] {
			c.JSON(200, gin.H{
				"code": 10009,
				"msg":  "token无效",
			})
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(403, gin.H{
				"code": 10087,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		//将当前userID保存到请求的上下文c中
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
