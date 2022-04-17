package router

import "net/http"

// api错误的结构体
type APIException struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Request string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, msg string) *APIException {
	return &APIException{
		Code:      code,
		Msg:       msg,
	}
}
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
