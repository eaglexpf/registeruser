package response

import (
	"github.com/go-playground/validator/v10"
	"registeruser/entity/response"
)

func (r *res) getParamError(err validator.ValidationErrors) map[string]string {
	result := make(map[string]string, 0)
	for _, v := range err {
		if field, ok := v.(validator.FieldError); ok {
			result[field.Field()] = field.Tag()
		}
	}
	return result
}

func (r *res) ErrorResponse(err error) *response.Response {
	result := Response.Code(503).Msg(err.Error()).Data(nil)
	valid_err, ok := err.(validator.ValidationErrors)
	switch true {
	case ok:
		result.Code(400).Msg("参数验证失败").Data(r.getParamError(valid_err))
	}
	return result.Result()
}
