package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
)

func GameList(ctx *gin.Context) {
	list, err := service.GameList()
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(list))
}
