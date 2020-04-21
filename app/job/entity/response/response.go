package response

import "registeruser/entity/response"

var Response *res

func init() {
	Response = &res{}
}

type res struct {
	code int64
	msg  string
	data interface{}
}

func (r *res) response() *response.Response {
	if r.data == nil {
		r.data = struct{}{}
	}
	return response.NewResponse(r.code, r.msg, r.data)
}

func (r *res) Code(code int64) *res {
	r.code = code
	return r
}

func (r *res) Msg(msg string) *res {
	r.msg = msg
	return r
}

func (r *res) Data(data interface{}) *res {
	r.data = data
	return r
}

func (r *res) Result() *response.Response {
	return r.response()
}

func (r *res) Success() *response.Response {
	r.code = 0
	r.msg = "success"
	r.data = struct{}{}
	return r.response()
}

func (r *res) SuccessMsg(msg string) *response.Response {
	r.msg = msg
	return r.response()
}

func (r *res) SuccessData(data interface{}) *response.Response {
	r.msg = "success"
	r.data = data
	return r.response()
}

func (r *res) SuccessMsgData(msg string, data interface{}) *response.Response {
	r.msg = msg
	r.data = data
	return r.response()
}

func (r *res) Error(code int64, msg string, data interface{}) *response.Response {
	r.code = code
	r.msg = msg
	r.data = data
	return r.response()
}

func (r *res) ErrorMsg(code int64, msg string) *response.Response {
	r.code = code
	r.msg = msg
	return r.response()
}

func (r *res) ErrorData(code int64, data interface{}) *response.Response {
	r.code = code
	r.data = data
	return r.response()
}
