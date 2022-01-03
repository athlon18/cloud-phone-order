package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func ExternalRouter(router *gin.RouterGroup) {
	router.POST("external/bind/:code", controller.BindMachine)                  // 绑定机器码和特征码
	router.GET("external/order/:code/new", controller.GetNewOrder)              // 获取新订单
	router.GET("external/order/:code/ing", controller.GetIngOrder)              // 获取执行中的订单
	router.POST("external/order/:code/update/:orderId", controller.UpdateOrder) // 更新订单
	router.POST("external/test/:code", controller.Test)                         // 测试
}

func ExternalRouterV2(router *gin.RouterGroup) {
	router.POST("external/login/register", controller.ExternalLoginRegister)            // 机器码登录绑定
	router.POST("external/machine/register", controller.ExternalRegister)               // 机器码注册
	router.POST("external/machine/deregister", controller.ExternalDeregister)           // 机器码注销
	router.POST("external/machine/health", controller.ExternalHealth)                   // 机器码心跳
	router.GET("external/order/:code/new", controller.ExternalGetNewOrder)              // 获取新订单
	router.GET("external/order/:code/ing", controller.ExternalGetIngOrder)              // 获取执行中的订单
	router.POST("external/order/:code/update/:orderId", controller.ExternalUpdateOrder) // 更新订单

}
