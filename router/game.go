package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func GameRouter(router *gin.RouterGroup)  {
	router.GET("game/list", controller.GameList)
}
