package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/model"
	"order/service"
)

func CheckLogin(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, "登录令牌已失效，请重新登录！", nil))
		ctx.Abort()
	}
	_, err := service.GetTokenInfo(authorization)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		ctx.Abort()
	}
	ctx.Next()
}

func GetLoginUser(ctx *gin.Context) (model.User, error) {
	user := model.User{}
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		return user, errors.New("登录令牌已失效，请重新登录！")
	}
	return service.GetTokenInfo(authorization)
}
