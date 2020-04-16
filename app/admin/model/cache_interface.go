package model

import "context"

type AdminUserCache interface {
	FindNameByUserUUID(ctx context.Context, uuid, name string) (ok bool, err error)
	SetUserServiceByUUID(ctx context.Context, uuid string, data []string)
}
