package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func InfoRouter(router *gin.RouterGroup)  {
	router.POST("user/info", controller.UserInfo)
	router.POST("user/logout", controller.UserLogout)
}
