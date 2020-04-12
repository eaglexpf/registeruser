// 实体数据 response
package response

import "net/http"

// 项目输出总格式
type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 新建一个Response数据
func NewResponse(code int64, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 服务器错误
func ErrorServiceUnavailable() *Response {
	return NewResponse(http.StatusServiceUnavailable, "服务器错误", nil)
}

// 权限验证失败
func ErrorForBidden() *Response {
	return NewResponse(http.StatusForbidden, "权限验证失败", nil)
}
