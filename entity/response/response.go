package response

import "net/http"

//type Response interface {
//}

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int64, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ErrorServiceUnavailable() *Response {
	return NewResponse(http.StatusServiceUnavailable, "服务器错误", nil)
}

func ErrorForBidden() *Response {
	return NewResponse(http.StatusForbidden, "权限验证失败", nil)
}
