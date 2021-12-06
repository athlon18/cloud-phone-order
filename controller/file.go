package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
)

func File(ctx *gin.Context) {
	file := service.GetFile()
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(file))
}
