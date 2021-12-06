package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.POST("register", controller.Register)
}
