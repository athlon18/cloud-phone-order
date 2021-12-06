package router

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func CategoriesRouter(router *gin.RouterGroup)  {
	router.GET("categories/list/:gameId", controller.CategoriesList)
}
