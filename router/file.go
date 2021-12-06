package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func FileRouter(router *gin.RouterGroup)  {
	router.GET("file", controller.File)
}
