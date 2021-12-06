package middleware

import (
	"github.com/gin-gonic/gin"
)

//跨域处理
func CrosMiddleware(Ctx *gin.Context) {
	origin := Ctx.Request.Header.Get("Origin")
	Ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	Ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	Ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Cookie, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie")
	Ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if Ctx.Request.Method == "OPTIONS" {
		Ctx.JSON(200, "Options Request!")
	}
	Ctx.Next()
}

