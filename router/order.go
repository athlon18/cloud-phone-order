package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func OrderRouter(router *gin.RouterGroup) {
	router.GET("order/list", controller.OrderList)
	router.POST("order/submit", controller.OrderSubmit)
	router.POST("order/change/:orderId", controller.OrderUpdateStatus)
}
