package auth

import (
	"context"
	"registeruser/app/admin/service"
	"registeruser/util/common"
)

var srv service.ServiceAuth

func init() {
	srv = service.NewServiceAuth()
}

func Check(ctx context.Context, name string, uuid string) {
	user, err := srv.FindUserByUUID(ctx, uuid)
	if err != nil {
		return
	}
	// 获取user_id的所有权限
	ids := srv.FindAllPermissionByUserID(ctx, user.ID)
	ids = common.SliceDuplicateInt64(ids)
	// 根据id查询名称别名
	// 判断服务是否在列表中
}
