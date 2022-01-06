package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func MachineRouter(router *gin.RouterGroup) {
	router.POST("machine/list", controller.MachineList)
	router.POST("machine/edit/:id", controller.Edit)
}
