package redis

import (
	"github.com/go-redis/redis"
	"registeruser/app/job/model"
	"registeruser/conf/global"
	"strconv"
	"time"
)

func NewSortSet() model.SortSet {
	return &sortSet{
		db:  global.REDIS,
		key: "task:sort_set",
	}
}

type sortSet struct {
	db  *redis.Client
	key string
}

func (s *sortSet) Push(uuid string, time_scores int64) (err error) {
	//err = s.db.Do("zadd", s.key, time_scores, uuid).Err()
	return s.db.ZAdd(s.key, redis.Z{
		Score:  float64(time_scores),
		Member: uuid,
	}).Err()
}

func (s *sortSet) Pull() (result []string, err error) {
	op := redis.ZRangeBy{
		Min: "1",
		Max: strconv.FormatInt(time.Now().Unix(), 10),
	}
	data, err := s.db.ZRangeByScore(s.key, op).Result()
	if err != nil {
		return
	}
	for _, v := range data {
		result = append(result, v)
	}
	return
	//data, err := s.db.Do("ZRANGEBYSCORE", s.key, 0, time.Now().Unix()).Result()
	//if err != nil {
	//	return
	//}
	//result_list, ok := data.([]interface{})
	//if !ok {
	//	return
	//}
	//for _, v := range result_list {
	//	if v_data, ok := v.(string); ok {
	//		if err := s.Remove(v_data); err == nil {
	//			result = append(result, v_data)
	//		}
	//	}
	//}
	//return
}

func (s *sortSet) Remove(uuid string) {
	//err := s.db.Do("ZREM", s.key, uuid).Err()
	//if err != nil {
	//	fmt.Println("有序集合删除失败", uuid)
	//}
	s.db.ZRem(s.key, uuid)
}
