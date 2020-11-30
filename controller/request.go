package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登陆")

//  GetCurrentUser  获取当前登陆的用户id
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {

	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}
