package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/middleware"
	"order/model"
	"order/service"
	"strconv"
)

type BindMachineExample struct {
	MachineCode string `json:"machine_code" example:"23333333333"`
}

// BindMachine
// @Tags V1接口
// @Summary 绑定特征码和机器码
// @Description 绑定特征码和机器码
// @Accept  json
// @Produce  json
// @Param code path string true  "特征码" default(PKPV48)
// @Param machine_code body BindMachineExample true "机器码"
// @Success 200 {object} object{success=bool,code=int}
// @Router /api/v1/external/bind/{code} [post]
func BindMachine(ctx *gin.Context) {
	user, err := middleware.GetCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if user.Machine != "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "账号已绑定机器码！", nil))
		return
	}
	data := struct {
		MachineCode string `json:"machine_code"`
	}{}
	if err = ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if data.MachineCode == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空！", nil))
		return
	}
	user.Machine = data.MachineCode
	if err = service.UpdateUserByCode(user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空！", nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(nil))

}

// GetNewOrder
// @Tags V1接口
// @Summary 获取一个新的订单
// @Description 获取一个新的订单，如果有正在执行的订单会返回报错
// @Accept  json
// @Produce  json
// @Param code path string true  "特征码" default(PKPV48)
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v1/external/order/{code}/new [get]
func GetNewOrder(ctx *gin.Context) {
	user, err := middleware.GetCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order, err := service.GetNewOrder(user)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order.Type = service.GetCategoryCode(order.Type)
	result := util.ToJsonMap(order)
	result["option"] = service.GetOptionSelectCodeByIds(order.Option)
	result["order_id"] = strconv.FormatInt(order.OrderId, 10)
	go service.CreateLog(order.OrderId, "订单开始执行", 0, "1")
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(result))
}

// GetIngOrder
// @Tags V1接口
// @Summary 获取执行中的订单
// @Description 获取执行中的订单，如果没有执行中的订单会返回报错
// @Accept  json
// @Produce  json
// @Param code path string true  "特征码" default(PKPV48)
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v1/external/order/{code}/ing [get]
func GetIngOrder(ctx *gin.Context) {
	user, err := middleware.GetCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order, err := service.GetIngOrder(user)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order.Type = service.GetCategoryCode(order.Type)
	result := util.ToJsonMap(order)
	result["option"] = service.GetOptionSelectCodeByIds(order.Option)
	result["order_id"] = strconv.FormatInt(order.OrderId, 10)
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(result))
}

type UpdateOrderExample struct {
	Status string `json:"status" example:"1"`
	Cnum   string `json:"cnum" example:"1"`
	Log    string `json:"log" example:"测试log"`
}

// UpdateOrder
// @Tags V1接口
// @Summary 更新的执行中的订单
// @Description 更新的执行中的订单 <br><br> status 状态（0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单）<br> cnum 完成数量 <br> text 日志内容
// @Accept  json
// @Produce  json
// @Param code path string true  "特征码" default(PKPV48)
// @Param order_id path int64 true  "订单ID" default(20210619215350)
// @Param body body UpdateOrderExample false "请求参数"
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v1/external/order/{code}/update/{order_id} [post]
func UpdateOrder(ctx *gin.Context) {
	_, err := middleware.GetCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	data := struct {
		Status string `json:"status"` // status 0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单
		CNum   string `json:"cnum"`   // 完成数量
		Log    string `json:"log"`    // 日志
	}{}
	if err = ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	if err = service.UpdateOrder(int64(orderId), data.Status, data.CNum); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if data.Log != "" {
		num, _ := strconv.Atoi(data.CNum)
		go service.CreateLog(int64(orderId), data.Log, num, data.Status)
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(strconv.Itoa(orderId)))
}

// Test
// @Tags V1接口
// @Summary 测试post接口
// @Description 测试post接口
// @Accept  json
// @Produce  json
// @Param code path string true  "特征码" default(PKPV48)
// @Success 200 {object} object{success=bool,data=model.User,code=int}
// @Router /api/v1/external/test/{code} [post]
func Test(ctx *gin.Context) {
	user, err := middleware.GetCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(user))

}

type ExternalLoginRegisterExample struct {
	MachineCode string `json:"machine_code" example:"23333333333"`
	Name        string `json:"username" example:"admin"`
	Password    string `json:"password" example:"qianyi"`
}

// ExternalLoginRegister
// @Tags V2接口
// @Summary 登录绑定机器码 返回用户机器码
// @Description 登录绑定机器码
// @Accept  json
// @Produce  json
// @Param machine_code body ExternalLoginRegisterExample true "机器码"
// @Success 200 {object} object{success=bool,code=int,message=string,data=model.UserMachine}
// @Router /api/v2/external/login/register [post]
func ExternalLoginRegister(ctx *gin.Context) {
	user := ExternalLoginRegisterExample{}
	_ = model.UserMachine{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	}

	if user.MachineCode == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空! ", nil))
	}

	userMachine, err := service.ExternalLoginRegister(user.Name, util.EncryptSha256(user.Password), user.MachineCode)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, util.Result().SetSuccess(userMachine))
}

type ExternalRegisterExample struct {
	MachineCode string `json:"machine_code" example:"23333333333"`
}

// ExternalRegister
// @Tags V2接口
// @Summary 机器服务发现
// @Description 服务发现(机器) 初始化时候调用
// @Accept  json
// @Produce  json
// @Param machine_code body ExternalRegisterExample true "机器码"
// @Success 200 {object} object{success=bool,code=int,message=string}
// @Router /api/v2/external/machine/register [post]
func ExternalRegister(ctx *gin.Context) {
	user := ExternalRegisterExample{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	}
	if user.MachineCode == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空! ", nil))
	}

	if err := service.ExternalRegister(user.MachineCode); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}

//ExternalDeregister
// @Tags V2接口
// @Summary 机器服务注销
// @Description 服务注销(机器) 关闭时候调用
// @Accept  json
// @Produce  json
// @Param machine_code body ExternalRegisterExample true "机器码"
// @Success 200 {object} object{success=bool,code=int,message=string}
// @Router /api/v2/external/machine/deregister [post]
func ExternalDeregister(ctx *gin.Context) {
	user := ExternalRegisterExample{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	}
	if user.MachineCode == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空! ", nil))
	}

	rows, err := service.ExternalDeregister(user.MachineCode)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if rows == 0 {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "未找到机器码，可能已经注销！", nil))
		return
	}

	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}

//ExternalHealth
// @Tags V2接口
// @Summary 机器心跳包
// @Description 机器心跳包
// @Accept  json
// @Produce  json
// @Param machine_code body ExternalRegisterExample true "机器码"
// @Success 200 {object} object{success=bool,code=int,message=string}
// @Router /api/v2/external/machine/health [post]
func ExternalHealth(ctx *gin.Context) {
	user := ExternalRegisterExample{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
	}
	if user.MachineCode == "" {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "机器码不能为空! ", nil))
	}

	rows, err := service.ExternalHealth(user.MachineCode)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if rows == 0 {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, "未找到机器码，可能已经注销！", nil))
		return
	}

	ctx.JSON(http.StatusOK, util.Result().SetSuccess(true))
}

// ExternalGetNewOrder
// @Tags V2接口
// @Summary 获取一个新的订单
// @Description 获取一个新的订单，如果有正在执行的订单会返回报错
// @Accept  json
// @Produce  json
// @Param code path string true  "用户机器码" default(085ebd3b-0c6d-469e-9cb5-f9971d81d223)
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v2/external/order/{code}/new [get]
func ExternalGetNewOrder(ctx *gin.Context) {
	machine, err := middleware.GetMachineCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order, err := service.GetMachineNewOrder(machine)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order.Type = service.GetCategoryCode(order.Type)
	result := util.ToJsonMap(order)
	result["option"] = service.GetOptionSelectCodeByIds(order.Option)
	result["order_id"] = strconv.FormatInt(order.OrderId, 10)
	go service.CreateLog(order.OrderId, "订单开始执行", 0, "1")
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(result))
}

// ExternalGetIngOrder
// @Tags V2接口
// @Summary 获取执行中的订单
// @Description 获取执行中的订单，如果没有执行中的订单会返回报错
// @Accept  json
// @Produce  json
// @Param code path string true  "用户机器码" default(085ebd3b-0c6d-469e-9cb5-f9971d81d223)
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v2/external/order/{code}/ing [get]
func ExternalGetIngOrder(ctx *gin.Context) {
	machine, err := middleware.GetMachineCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order, err := service.GetMachineIngOrder(machine)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	order.Type = service.GetCategoryCode(order.Type)
	result := util.ToJsonMap(order)
	result["option"] = service.GetOptionSelectCodeByIds(order.Option)
	result["order_id"] = strconv.FormatInt(order.OrderId, 10)
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(result))
}

// ExternalUpdateOrder
// @Tags V2接口
// @Summary 更新的执行中的订单
// @Description 更新的执行中的订单 <br><br> status 状态（0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单）<br> cnum 完成数量 <br> text 日志内容
// @Accept  json
// @Produce  json
// @Param code path string true  "用户机器码" default(085ebd3b-0c6d-469e-9cb5-f9971d81d223)
// @Param order_id path int64 true  "订单ID" default(20210619215350)
// @Param body body UpdateOrderExample false "请求参数"
// @Success 200 {object} object{success=bool,data=model.Order,code=int}
// @Router /api/v2/external/order/{code}/update/{order_id} [post]
func ExternalUpdateOrder(ctx *gin.Context) {
	_, err := middleware.GetMachineCode(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	data := struct {
		Status string `json:"status"` // status 0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单
		CNum   string `json:"cnum"`   // 完成数量
		Log    string `json:"log"`    // 日志
	}{}
	if err = ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	if err = service.UpdateOrder(int64(orderId), data.Status, data.CNum); err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if data.Log != "" {
		num, _ := strconv.Atoi(data.CNum)
		go service.CreateLog(int64(orderId), data.Log, num, data.Status)
	}
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(strconv.Itoa(orderId)))
}
