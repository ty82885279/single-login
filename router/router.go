package router

import (
	"github.com/gin-gonic/gin"
	"single-login/controller"
	"single-login/middleware"
)

func Setup() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", controller.Login)
		v1.POST("/alisa", controller.AlisaHandler)
		v1.POST("/token", controller.RefreshTokenHandler)
	}
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.Post)
		v1.POST("/logout", controller.UserLogOut)

	}

	return r
}
