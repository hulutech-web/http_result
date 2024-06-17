package http_result

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

/*
*
Success(200,"操作成功",data)
Error(500,"服务器错误",data)
*/
type Instance interface {
	Success(message string, data interface{}) http.Response
	Error(code int, message string, data interface{}) http.Response
	ValidError(message string, errors map[string]map[string]string) http.Response
}

type HttpResult struct {
	ctx      http.Context
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Response http.Response
}

func NewResult(ctx http.Context) *HttpResult {
	return &HttpResult{
		ctx: ctx,
	}
}

/*
Success 默认成功返回
*/
func (h *HttpResult) Success(message string, data interface{}) http.Response {
	if message == "" {
		message = facades.Config().GetString("http_result.Message")
	}
	return h.ctx.Response().Success().Json(http.Json{
		"message": message,
		"data":    data,
	})
}

/*
*
Error
自定义错误
*/
func (h *HttpResult) Error(code int, message string, data interface{}) http.Response {
	if message == "" {
		message = facades.Config().GetString("http_result.Message")
	}
	if code == 0 {
		code = facades.Config().GetInt("http_result.Code")
	}
	h.ctx.Request().AbortWithStatusJson(code, http.Json{
		"message": message,
		"data":    data,
	})
	return nil
}

/*
*
ValidError
表单验证服务
*/
func (h *HttpResult) ValidError(message string, errors map[string]map[string]string) http.Response {
	h.Code = 422
	if message == "" {
		message = facades.Config().GetString("http_result.Message", "验证失败")
	}
	h.ctx.Request().AbortWithStatusJson(h.Code, http.Json{
		"message": message,
		"errors":  errors,
	})
	return nil
}
