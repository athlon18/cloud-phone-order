package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/middleware"
	"order/model"
	"order/service"
	"strconv"
	"strings"
)

func OrderList(ctx *gin.Context) {
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	status := ctx.DefaultQuery("status", "")
	orderId, _ := strconv.Atoi(ctx.DefaultQuery("order_id", "0"))
	name := ctx.DefaultQuery("name", "")
	gameId, _ := strconv.Atoi(ctx.DefaultQuery("order_id", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	data, total, _ := service.OrderList(user.ID, name, uint(gameId), int64(orderId), status, pageSize, pageNo)
	for index, item := range data {
		data[index].Type = service.GetCategoryName(item.Type)
	}
	//for i := 0; i < 200; i++ {
	//	data = append(data, data[0])
	//}
	ctx.JSON(http.StatusOK, util.Result().SetTotal(int(total)).SetSuccess(data))
}
func OrderSubmit(ctx *gin.Context) {
	user, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	data := struct {
		GameId   uint           `json:"game_id"`
		ModeId   uint           `json:"mode_id"`
		Name     string         `json:"name"`
		Password string         `json:"password"`
		Num      int            `json:"num"`
		Option   map[string]int `json:"option"`
		Type     []string       `json:"type"`
		Machine  string         `json:"machine"`
	}{}
	if err = ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if data.Machine == "-1" {
		machine, u := service.GetUserMachine(user.ID)
		if u != nil {
			ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
			return
		}
		data.Machine = machine.Machine
	}

	option, _ := json.Marshal(data.Option)
	orderId := util.GetOrderID()
	if err = service.OrderSubmit(&model.Order{
		OrderId:  orderId,
		GameId:   data.GameId,
		ModeId:   data.ModeId,
		UserId:   user.ID,
		Name:     data.Name,
		Password: data.Password,
		Type:     strings.Join(data.Type, ","),
		Option:   string(option),
		Num:      data.Num,
		CNum:     0,
		Machine:  data.Machine,
	}); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	go service.CreateLog(orderId, "订单初始化", 0, "0")
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}

func OrderUpdateStatus(ctx *gin.Context) {
	_, err := middleware.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusNotFound, err.Error(), nil))
		return
	}
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	data := struct {
		Status int `json:"status"`
	}{}
	if err = ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order, err := service.UpdateOrderByStatus(int64(orderId), data.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	go service.CreateLog(int64(orderId), "用户更新订单状态", order.CNum, strconv.Itoa(data.Status))
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}
