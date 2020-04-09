package app

import (
	"github.com/go-redis/redis"
	"registeruser/entity/global"
	"registeruser/log"
)

func InitRedis() {
	config := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Errorf("redis connect err:%v", err)
	} else {
		global.REDIS = client
		log.Info("redis已链接...")
	}
}
