package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/controller"
)

func UploadRouter(router *gin.RouterGroup) {
	router.GET("html/index", controller.UploadHtml)
	router.POST("upload", controller.Upload)
	router.StaticFS("upload", http.Dir("./upload"))
}
