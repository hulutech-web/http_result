package contracts

import "github.com/goravel/framework/contracts/http"

type HttpResult interface {
	Success(message string, data interface{}) http.Response
	Error(code int, message string, data interface{}) http.Response
	ValidError(message string, errors map[string]map[string]string) http.Response
}
