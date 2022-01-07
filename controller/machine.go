package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/middleware"
	"order/model"
	"order/service"
)

func MachineList(ctx *gin.Context) {
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	body := model.UserMachine{}
	if err = ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	machineList, err := service.UserMachineList(body, user.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(machineList))
}

func Edit(ctx *gin.Context) {
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	id := ctx.Param("id")
	body := model.UserMachine{}
	if err = ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	if rows := service.EditUserMachine(body, id, user.ID); rows == 0 {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "更新失败！", nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}
