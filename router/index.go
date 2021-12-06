package router

import (
	"github.com/gin-gonic/gin"
	_ "order/docs"
	"order/middleware"
)

var swagHandler gin.HandlerFunc

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := gin.Default()
	router.LoadHTMLFiles("html/upload.html")
	router.MaxMultipartMemory = 256

	// 中间件拦截
	router.Use(func(c *gin.Context) {
		//跨域设置
		middleware.CrosMiddleware(c)
	})

	// 开启swag
	if swagHandler != nil {
		router.GET("/swagger/*any", swagHandler)
	}

	//api
	api := router.Group("api/")
	{
		IndexRouter(api)
		IndexV2Router(api)
	}
	return router
}

func IndexRouter(api *gin.RouterGroup) {
	// V1版本
	v1 := api.Group("v1/")
	LoginRouter(v1)
	RegisterRouter(v1)
	InfoRouter(v1)
	GameRouter(v1)
	CategoriesRouter(v1)
	OrderRouter(v1)
	ExternalRouter(v1)
	UploadRouter(v1)
	FileRouter(v1)
}

func IndexV2Router(api *gin.RouterGroup) {
	// V2版本
	v2 := api.Group("v2/")
	ExternalRouterV2(v2)
}
