// 返回数据定义
package response

import (
	app_response "registeruser/entity/response"
)

// 返回数据根节点
type Response struct {
	*app_response.Response
}

type ResponsePage struct {
	Meta *ResponsePageMeta `json:"meta"`
	List interface{}       `json:"list"`
}
type ResponsePageMeta struct {
	Count     int64 `json:"count"`
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	PageCount int64 `json:"page_count"`
}

//type Response struct {
//	Code int64       `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}

// 新建一个返回数据
func newResponse(code int64, msg string, data interface{}) *Response {
	return &Response{
		app_response.NewResponse(code, msg, data),
	}
	//return &Response{
	//	Code: code,
	//	Msg:  msg,
	//	Data: data,
	//}
}

// 新建一个返回数据，默认data为nil
func newResponseNilData(code int64, msg string) *Response {
	return &Response{
		app_response.NewResponse(code, msg, nil),
	}
	//return &Response{
	//	Code: code,
	//	Msg:  msg,
	//	Data: nil,
	//}
}

// 失败；自定义类型和数据
func Error(code int64, msg string) *Response {
	return newResponseNilData(code, msg)
}

// 失败：自定义类型，消息和数据
func ErrorData(code int64, msg string, data interface{}) *Response {
	return newResponse(code, msg, data)
}

// 请求参数验证失败
func ErrorParamValidate() *Response {
	return newResponseNilData(400, "请求参数参悟")
}

// 请求参数验证失败，自定义消息
func ErrorParamValidateMsg(msg string) *Response {
	return newResponseNilData(400, msg)
}

// 请求参数验证失败：自定义数据
func ErrorParamValidateData(data interface{}) *Response {
	return newResponse(400, "参数错误", data)
}

// 请求参数验证失败：自定义消息和数据
func ErrorParamValidateMsgData(msg string, data interface{}) *Response {
	return newResponse(400, msg, data)
}
