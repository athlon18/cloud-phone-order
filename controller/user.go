package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/middleware"
	"order/service"
)

func UserInfo(ctx *gin.Context) {
	userInfo := struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
		Code         string   `json:"code"`
	}{
		Roles:  []string{"editor"},
		Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	userInfo.Name = user.Name
	userInfo.Code = user.Code
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(userInfo))
}

func UserLogout(ctx *gin.Context) {
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	err = service.UpdateUserTokenByID(int(user.ID))
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(err))
}
