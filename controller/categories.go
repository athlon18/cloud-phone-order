package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
	"strconv"
)

func CategoriesList(ctx *gin.Context) {
	gameId, _ := strconv.Atoi(ctx.Param("gameId"))
	data := service.GetCategoriesList(gameId)

	ctx.JSON(http.StatusOK, util.Result().SetSuccess(data))
}
