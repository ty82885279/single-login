package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	{
		"code":10001, //错误码
		"msg":"XXX",  //提示信息
		"data":{}     //数据
	}
*/
type ResponseData struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1001,
		Msg:  "请求成功",
		Data: data,
	})
}

func ResponseSuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1001,
		Msg:  msg,
		Data: data,
	})
}
func ResponseError(c *gin.Context) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1002,
		Msg:  "token 错误",
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 1003,
		Msg:  msg,
	})
}
