package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
)

func Register(ctx *gin.Context) {
	register := struct {
		Name            string `form:"username" json:"username" binding:"required,min=4"`
		Password        string `form:"password" json:"password" binding:"required,min=4"`
		ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password" json:"confirm_password"`
	}{}
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	user, err := service.Register(register.Name, register.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(user))
}
