package http_result

import (
	"github.com/goravel/framework/contracts/database/orm"
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
	SearchByParams(params map[string]string, excepts ...string) *HttpResult
	ResultPagination(dest any) (http.Response, error)
}

type HttpResult struct {
	Query    orm.Query
	Context  http.Context
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Response http.Response
}

func NewResult(ctx http.Context) *HttpResult {
	return &HttpResult{
		Context: ctx,
		Query:   facades.Orm().Query(),
	}
}

/*
Success 默认成功返回
*/
func (h *HttpResult) Success(message string, data interface{}) http.Response {
	if message == "" {
		message = facades.Config().GetString("http_result.Message")
	}
	//查询的结果有可能存在data为nil的情况，判断如果是nil则，给定一个[]string{}类型的数据，避免前端页面报错
	//或者判断data的slice类型的长度为0
	res := []string{}
	if data == nil {
		return h.Context.Response().Success().Json(http.Json{
			"message": message,
			"data":    res,
		})
	}
	slice, ok := data.([]interface{})
	if ok && len(slice) == 0 {
		return h.Context.Response().Success().Json(http.Json{
			"message": message,
			"data":    res,
		})
	}

	return h.Context.Response().Success().Json(http.Json{
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
	h.Context.Request().AbortWithStatusJson(code, http.Json{
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
	h.Context.Request().AbortWithStatusJson(h.Code, http.Json{
		"message": message,
		"errors":  errors,
	})
	return nil
}
