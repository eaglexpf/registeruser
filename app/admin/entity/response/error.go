package response

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	// 用户权限检查
	ERROR_AUTH = errors.New("权限验证失败")

	// param参数验证失败
	ERROR_PARAM_VALIDATE = errors.New("参数验证失败")

	// 查询用户基本信息
	ERROR_USER_FIND_BY_EMAIL = errors.New("无效的邮箱")

	// 用户注册
	ERROR_USER_UNIQUE_NAME    = errors.New("账号已存在")
	ERROR_USER_UNIQUE_EMAIL   = errors.New("邮箱已存在")
	ERROR_USER_PASSWORD_EQUAL = errors.New("两次密码不一致")
	ERROR_USER_PASSWORD_CRYPT = errors.New("密码加密失败")
	ERROR_USER_REGISTER_DB    = errors.New("用户注册失败")

	// 用户登录
	ERROR_USER_PASSWORD_FALSE = errors.New("密码错误")
	ERROR_USER_FIND_BY_NAME   = errors.New("账号不存在")
	ERROR_JWT_CREATE          = errors.New("token生成失败")

	// 修改用户基本信息
	ERROR_USER_FIND_BY_UUID   = errors.New("无效的UUID")
	ERROR_USER_UPDATE_INFO_DB = errors.New("修改失败")

	// 修改用户密码
	ERROR_USER_UPDATE_PASSWORD_DB = errors.New("修改失败")

	// 角色列表查询
	ERROR_ROLE_LIST       = errors.New("查询失败")
	ERROR_ROLE_FIND_BY_ID = errors.New("无效的角色")
	ERROR_ROLE_REGISTER   = errors.New("添加失败")
	ERROR_ROLE_UPDATE     = errors.New("修改失败")
	ERROR_ROLE_DELETE     = errors.New("删除失败")

	// 权限分配
	ERROR_PERMISSION_LIST               = errors.New("查询失败")
	ERROR_PERMISSION_REGISTER           = errors.New("添加失败")
	ERROR_PERMISSION_DELETE             = errors.New("删除失败")
	ERROR_PERMISSION_DAO_PARENT_TYPE    = errors.New("parent_type类型只能为【user,role】")
	ERROR_PERMISSION_DAO_CHILDREN_TYPE  = errors.New("children_type类型只能为【role,api,service】")
	ERROR_PERMISSION_NONE_PARENT_TYPE   = errors.New("父权限类型不能为空")
	ERROR_PERMISSION_NONE_PARENT_ID     = errors.New("父权限id不能为空")
	ERROR_PERMISSION_NONE_CHILDREN_TYPE = errors.New("子权限类型不能为空")
	ERROR_PERMISSION_NONE_CHILDREN_ID   = errors.New("子权限id不能为空")
	ERROR_PERMISSION_NONE               = errors.New("参数不能为空")

	// 服务列表查询
	ERROR_SERVICE_LIST           = errors.New("查询失败")
	ERROR_SERVICE_UNIQUE_NAME    = errors.New("账号已存在")
	ERROR_SERVICE_FIND_BY_ID     = errors.New("无效的服务")
	ERROR_SERVICE_REGISTER       = errors.New("添加失败")
	ERROR_SERVICE_UPDATE         = errors.New("修改失败")
	ERROR_SERVICE_DELETE_BY_ID   = errors.New("删除失败")
	ERROR_SERVICE_DELETE_BY_NAME = errors.New("删除失败")
)

func getParamError(err validator.ValidationErrors) map[string]string {
	result := make(map[string]string, 0)
	for _, v := range err {
		if field, ok := v.(validator.FieldError); ok {
			result[field.Field()] = field.Tag()
		}
	}
	return result
}

func GetErrorResponse(err error) *Response {
	var default_code int64 = 503
	result := newResponse(default_code, err.Error(), struct{}{})
	valid_err, ok := err.(validator.ValidationErrors)
	switch true {
	case ok:
		result.Code = 400
		result.Msg = "参数验证失败"
		result.Data = getParamError(valid_err)
	case errors.Is(ERROR_PARAM_VALIDATE, err):
		result.Code = 400
	case errors.Is(ERROR_AUTH, err):
		result.Code = 1001
	case errors.Is(ERROR_USER_FIND_BY_EMAIL, err):
		// 无效的邮箱
		result.Code = 1011
	case errors.Is(ERROR_USER_UNIQUE_NAME, err):
		// 账号已存在
		result.Code = 1012
	case errors.Is(ERROR_USER_UNIQUE_EMAIL, err):
		// 邮箱已存在
		result.Code = 1013
	case errors.Is(ERROR_USER_PASSWORD_EQUAL, err):
		// 两次密码不一致
		result.Code = 1014
	case errors.Is(ERROR_USER_PASSWORD_CRYPT, err):
		// 密码加密失败
		result.Code = 1015
	case errors.Is(ERROR_USER_REGISTER_DB, err):
		// 用户注册失败
		result.Code = 1016
	case errors.Is(ERROR_USER_PASSWORD_FALSE, err):
		// 密码错误
		result.Code = 1021
	case errors.Is(ERROR_USER_FIND_BY_NAME, err):
		// 账号不存在
		result.Code = 1022
	case errors.Is(ERROR_JWT_CREATE, err):
		// token生成失败
		result.Code = 1023
	case errors.Is(ERROR_USER_FIND_BY_UUID, err):
		// 无效的UUID
		result.Code = 1031
	case errors.Is(ERROR_USER_UPDATE_INFO_DB, err):
		// 修改失败
		result.Code = 1032
	case errors.Is(ERROR_USER_UPDATE_PASSWORD_DB, err):
		// 修改失败
		result.Code = 1041
	case errors.Is(ERROR_ROLE_LIST, err):
		// 查询失败
		result.Code = 1051
	case errors.Is(ERROR_ROLE_FIND_BY_ID, err):
		// 无效的角色
		result.Code = 1052
	case errors.Is(ERROR_ROLE_REGISTER, err):
		// 添加失败
		result.Code = 1053
	case errors.Is(ERROR_ROLE_UPDATE, err):
		// 修改失败
		result.Code = 1054
	case errors.Is(ERROR_ROLE_DELETE, err):
		// 删除失败
		result.Code = 1055
	case errors.Is(ERROR_PERMISSION_LIST, err):
		// 查询失败
		result.Code = 1061
	case errors.Is(ERROR_PERMISSION_REGISTER, err):
		// 添加失败
		result.Code = 1062
	case errors.Is(ERROR_PERMISSION_DELETE, err):
		// 删除失败
		result.Code = 1063
	case errors.Is(ERROR_PERMISSION_DAO_PARENT_TYPE, err):
		// parent_type类型只能为【user,role】
		result.Code = 1064
	case errors.Is(ERROR_PERMISSION_DAO_CHILDREN_TYPE, err):
		// children_type类型只能为【role,api,service】
		result.Code = 1065
	case errors.Is(ERROR_PERMISSION_NONE_PARENT_TYPE, err):
		// 父权限类型不能为空
		result.Code = 1066
	case errors.Is(ERROR_PERMISSION_NONE_PARENT_ID, err):
		// 父权限id不能为空
		result.Code = 1068
	case errors.Is(ERROR_PERMISSION_NONE_CHILDREN_TYPE, err):
		// 子权限类型不能为空
		result.Code = 1069
	case errors.Is(ERROR_PERMISSION_NONE_CHILDREN_ID, err):
		// 子权限id不能为空
		result.Code = 1070
	case errors.Is(ERROR_PERMISSION_NONE, err):
		// 参数不能为空
		result.Code = 1071
	case errors.Is(ERROR_SERVICE_LIST, err):
		// 查询失败
		result.Code = 1081
	case errors.Is(ERROR_SERVICE_UNIQUE_NAME, err):
		// 账号已存在
		result.Code = 1082
	case errors.Is(ERROR_SERVICE_FIND_BY_ID, err):
		// 无效的服务
		result.Code = 1083
	case errors.Is(ERROR_SERVICE_REGISTER, err):
		// 添加失败
		result.Code = 1084
	case errors.Is(ERROR_SERVICE_UPDATE, err):
		// 修改失败
		result.Code = 1085
	case errors.Is(ERROR_SERVICE_DELETE_BY_ID, err):
		// 删除失败
		result.Code = 1086
	case errors.Is(ERROR_SERVICE_DELETE_BY_NAME, err):
		// 删除失败
		result.Code = 1087
	default:
		result.Code = default_code
	}
	return result
}
