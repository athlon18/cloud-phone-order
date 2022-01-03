package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"order/model"
	"order/service"
)

func GetCode(ctx *gin.Context) (user model.User, err error) {
	code := ctx.Param("code")
	if code == "" {
		return user, errors.New("特征码不存在，请重试！")
	}
	return service.GetCodeInfo(code)
}

func GetMachineCode(ctx *gin.Context) (user model.UserMachine, err error) {
	code := ctx.Param("code")
	if code == "" {
		return user, errors.New("机器码不存在，请重新登录！")
	}
	return service.GetMachineInfo(code)

}
