package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"single-login/dao/redis"
	"single-login/logic"
	"single-login/model"
	"single-login/pkg/jwt"
	"strconv"
	"strings"
)

func Login(c *gin.Context) {

	var user = new(model.User)
	err := c.ShouldBind(user)
	if err != nil {
		zap.L().Error("should bind err", zap.Error(err))
		fmt.Printf("should bind err:%v\n", err)
		return
	}

	UserIDStr, aToken, rToken, err1 := logic.FindUser(user)
	if err1 != nil {

		if errors.Is(err1, gorm.ErrRecordNotFound) {
			fmt.Printf("用户不存在")
			c.JSON(200, gin.H{
				"code":   1001,
				"msg":    "登陆错误",
				"detail": "用户名或密码错误",
			})
			return
		}

		return
	}
	c.JSON(200, gin.H{
		"code":         1000,
		"msg":          "登陆成功",
		"userID":       UserIDStr,
		"AccseeToken":  aToken,
		"RefreshToken": rToken,
	})

}
func UserLogOut(c *gin.Context) {
	userID := c.PostForm("userID")
	err := redis.UserLogOut(userID)
	if err != nil {
		zap.L().Error("UserLogOut err", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{
		"code":   1000,
		"msg":    "退出成功",
		"userID": userID,
	})
}

// 保存alisa
func AlisaHandler(c *gin.Context) {

	userID := c.PostForm("userID")
	alisa := c.PostForm("alisa")

	redis.SaveAlisa(userID, alisa)

}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	rToken := c.PostForm("refresh_token")
	userID := c.PostForm("userID")
	aToken := c.Request.Header.Get("Authorization")
	if aToken == "" {
		ResponseErrorWithMsg(c, "请求头中Token格式错误")
		zap.L().Error("请求头中Token格式错误")
		return
	}
	//按空格分割
	parts := strings.SplitN(aToken, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, &ResponseData{
			Code: 1002,
			Msg:  "token 错误",
			Data: nil,
		})
		return
	}
	fmt.Printf("参数----%s\n", rToken)
	Rtoken, err := redis.GetRToken(userID)
	if err != nil {

		c.JSON(200, gin.H{
			"code": 10087,
			"msg":  "Rtoken过期",
		})
		return
	}
	if rToken == Rtoken {

		uID, _ := strconv.ParseInt(userID, 10, 64)
		newAToken, newRToken := jwt.RefreshToken(uID)
		_ = redis.SaveToken(uID, newAToken, newRToken)
		if redis.UpdateAlisaTime(userID) && redis.UpdateLoginTime(userID) {
			c.JSON(http.StatusOK, gin.H{
				"code":          10000,
				"access_token":  newAToken,
				"refresh_token": newRToken,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 10002,
				"msg":  "服务器异常",
			})
		}
	}
}

// 发布
func Post(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 1000,
		"msg":  "发布成功",
	})
}
