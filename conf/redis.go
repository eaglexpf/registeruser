package conf

import (
	"github.com/go-redis/redis"
	"registeruser/conf/global"
	"registeruser/conf/log"
)

// 初始化redis链接
func initRedis() {
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

// 清理redis链接
func closeRedis() {
	if global.REDIS == nil {
		return
	}
	err := global.REDIS.Close()
	if err != nil {
		log.Infof("Redis关闭连接失败：%v", err)
	}
}
