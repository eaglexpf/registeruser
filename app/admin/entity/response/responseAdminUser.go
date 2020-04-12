package response

func ErrorRegisterAdminUser() *Response {
	return newResponseNilData(1001, "用户注册失败")
}

func ErrorFindAdminUserByUsername() *Response {
	return newResponseNilData(1002, "无效的账号")
}
