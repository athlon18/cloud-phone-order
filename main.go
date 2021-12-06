package main

import (
	"github.com/gin-gonic/gin"
	routers "order/router"
)
// @title 下单API
// @version 1.0
// @description 下单API
func main() {
	router := routers.InitRouter()
	run(router)
}

func run(router *gin.Engine) {

	_ = router.Run(":8000")
}
