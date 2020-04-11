package response

import (
	app_response "registeruser/entity/response"
)

type Response struct {
	*app_response.Response
}

//type Response struct {
//	Code int64       `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}

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

func Success(data interface{}) *Response {
	return newResponse(0, "Success", data)
}

func Error(code int64, msg string) *Response {
	return newResponseNilData(code, msg)
}

func ErrorData(code int64, msg string, data interface{}) *Response {
	return newResponse(code, msg, data)
}

func ErrorParamValidate() *Response {
	return newResponseNilData(400, "请求参数参悟")
}

func ErrorParamValidateMsg(msg string) *Response {
	return newResponseNilData(400, msg)
}

func ErrorParamValidateData(data interface{}) *Response {
	return newResponse(400, "参数错误", data)
}

func ErrorParamValidateMsgData(msg string, data interface{}) *Response {
	return newResponse(400, msg, data)
}
