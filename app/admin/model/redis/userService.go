package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"time"
)

const (
	USER_SERVICE_KEY = "user:%s:service-key"
)

func NewAdminUserCache() model.AdminUserCache {
	return &userServiceCache{db: global.REDIS}
}

type userServiceCache struct {
	db *redis.Client
}

func (u *userServiceCache) FindNameByUserUUID(ctx context.Context, uuid, name string) (ok bool, err error) {
	key := fmt.Sprintf(USER_SERVICE_KEY, uuid)
	if u.db.Exists(key).Val() == 0 {
		err = errors.New("key不存在")
		return
	}
	ok = u.db.SIsMember(key, name).Val()
	return
}

func (u *userServiceCache) SetUserServiceByUUID(ctx context.Context, uuid string, data []string) {
	key := fmt.Sprintf(USER_SERVICE_KEY, uuid)
	if u.db.Exists(key).Val() != 0 {
		u.db.Del(key)
	}
	data_inter := make([]interface{}, len(data))
	for k, v := range data {
		data_inter[k] = v
	}
	u.db.SAdd(key, data_inter...)
	u.db.Expire(key, time.Duration(global.CONFIG.Redis.ExpireUserService)*time.Second)
}
