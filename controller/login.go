package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
)

func Login(ctx *gin.Context) {
	user := struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	}
	data, err := service.Login(user.Name, util.EncryptSha256(user.Password))
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	} else {
		ctx.JSON(http.StatusOK, util.Result().SetSuccess(data))
	}
}
