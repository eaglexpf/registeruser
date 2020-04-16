package response

// 成功：自定义返回数据
func Success(data interface{}) *Response {
	if data == nil {
		data = struct{}{}
	}
	return newResponse(0, "Success", data)
}

// 成功：自定义返回数据
func Success1(msg string, data interface{}) *Response {
	if data == nil {
		data = struct{}{}
	}
	return newResponse(0, msg, data)
}

// 成功：自定义返回数据
func SuccessMsg(msg string) *Response {
	return newResponse(0, msg, struct{}{})
}

// 成功：自定义返回数据
func SuccessData(data interface{}) *Response {
	if data == nil {
		data = struct{}{}
	}
	return newResponse(0, "Success", data)
}
