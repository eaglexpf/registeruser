package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"registeruser/app/job/model"
	"registeruser/conf/global"
)

func NewList() model.List {
	return &list{
		db:  global.REDIS,
		key: "task:list",
	}
}

type list struct {
	db  *redis.Client
	key string
}

func (l *list) Push(uuid string) error {
	return l.db.RPush(l.key, uuid).Err()
	//return l.db.Do("RPUSH", l.key, uuid).Err()
}

func (l *list) Pull() (string, error) {
	// 无数据时 阻塞7200秒
	list, err := l.db.BLPop(0, l.key).Result()
	//fmt.Println("list pop 读取：", list, err, reflect.TypeOf(list))
	if err != nil {
		return "", err
	}
	if len(list) != 2 {
		return "", errors.New(fmt.Sprintf("查询到的数据为：%v", list))
	}
	return list[1], nil
	//data, err := l.db.Do("BLPOP", l.key, 0).Result()
	//fmt.Println("list pop 读取：", data, err)
	//if err != nil {
	//	return "", err
	//}
	//redis_data, ok := data.([]interface{})
	//if !ok {
	//	return "", errors.New("返回结果不是string类型")
	//}
	//if len(redis_data) != 2 {
	//	return "", errors.New(fmt.Sprintf("查询到的数据为：%v", data))
	//}
	//result, ok := redis_data[1].(string)
	//if !ok {
	//	return "", errors.New("返回结果不是string类型123")
	//}
	//return result, nil
}
