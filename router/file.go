package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func FileRouter(router *gin.RouterGroup) {
	router.GET("file", controller.File)
	router.GET("file/index", controller.FileHtml)
	router.POST("folder/create", controller.FolderCreate)
	router.POST("file/delete", controller.FileDelete)
	router.GET("file/delete/:id", controller.UploadDelete)
}
