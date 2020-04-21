package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"registeruser/app/job/model"
	"registeruser/conf/global"
)

func NewSet() model.Set {
	return &set{
		db:  global.REDIS,
		key: "task:set",
	}
}

type set struct {
	db  *redis.Client
	key string
}

// 向集合中添加数据
func (s *set) Push(uuid string) error {
	return s.db.SAdd(s.key, uuid).Err()
}

// 判断数据是否存在
func (s *set) Exist(uuid string) (bool, error) {
	return s.db.SIsMember(s.key, uuid).Result()
}

// 从集合中移除数据
func (s *set) Remove(uuid string) (is bool, err error) {
	is, err = s.db.SIsMember(s.key, uuid).Result()
	if err != nil {
		return
	}
	if is {
		err = s.db.SRem(s.key, uuid).Err()
	}
	return
}

func (s *set) Pop() (string, error) {
	num, err := s.db.SCard(s.key).Result()
	if err != nil {
		return "", err
	}
	if num < 1 {
		return "", fmt.Errorf("集合中只有%d条数据了", num)
	}
	data, err := s.db.SPop(s.key).Result()
	if err != nil {
		return "", err
	}
	return data, nil
	//num, err := s.db.Do("SCARD", s.key).Int()
	//if err != nil {
	//	return "", err
	//}
	//if num < 1 {
	//	return "", fmt.Errorf("集合中只有%d条数据了", num)
	//}
	//data, err := s.db.Do("SPOP", s.key).Result()
	//if err != nil {
	//	return "", err
	//}
	//result, ok := data.(string)
	//if !ok {
	//	return "", fmt.Errorf("集合读取数据格式失败：%v", data)
	//}
	//return result, nil
}
