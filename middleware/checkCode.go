package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"order/model"
	"order/service"
)

func GetCode(ctx *gin.Context) (model.User, error) {
	user := model.User{}
	code := ctx.Param("code")
	if code == "" {
		return user, errors.New("特征码不存在，请重试！")
	}
	return service.GetCodeInfo(code)
}
