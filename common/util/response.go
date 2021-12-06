package util

import "encoding/json"

type ResultBean struct {
	IsSuccess   bool        `json:"success"`           //是否成功 系统错误或业务错误都返回false
	Code        int         `json:"code"`              //错误代码
	ErrorMsg    string      `json:"message,omitempty"` //报错详细信息 注：只针对于系统运行错误，业务代码错误一律只返回编码后再用多语言做处理
	ResultData  interface{} `json:"data,omitempty"`    //返回的json数据字符串
	ResultTotal int         `json:"total,omitempty"`   //查询的总数
}

func Result() *ResultBean {
	return &ResultBean{}
}

// 处理成功
func (this *ResultBean) SetSuccess(resultData interface{}) *ResultBean {
	this.IsSuccess = true
	this.Code = 200
	this.ResultData = resultData
	return this
}

// 处理成功
func (this *ResultBean) SetTotal(total int) *ResultBean{
	this.ResultTotal = total
	return this
}

// 处理失败
func (this *ResultBean) SetError(errCode int, errMsg string, resultData interface{}) *ResultBean {
	this.IsSuccess = false
	this.Code = errCode
	this.ErrorMsg = errMsg
	this.ResultData = resultData
	return this
}

//将本对象转为json
func (this *ResultBean) ToJson() string {
	str, _ := json.Marshal(this)
	return string(str)
}
