package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func LoginRouter(router *gin.RouterGroup) {
	router.POST("/login", controller.Login)
}
