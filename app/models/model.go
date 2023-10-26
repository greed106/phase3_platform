package models

import "net/http"

// Result 给前端的统一模型
type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// NewResultSuccessWithoutData 创建一个成功的响应，但不带任何数据
func NewResultSuccessWithoutData() *Result {
	return &Result{Code: http.StatusOK}
}

// NewResultSuccessWithData 创建一个成功的响应，并带有数据
func NewResultSuccessWithData(data interface{}) *Result {
	return &Result{Code: http.StatusOK, Data: data}
}

// NewResultError 创建一个失败的响应
func NewResultError(msg string) *Result {
	return &Result{Code: http.StatusBadRequest, Msg: msg}
}
