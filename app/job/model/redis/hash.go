package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"registeruser/app/job/entity/dao"
	"registeruser/app/job/model"
	"registeruser/conf/global"
)

func NewHash() model.Hash {
	return &hash{
		db:  global.REDIS,
		key: "task:hash",
	}
}

type hash struct {
	db  *redis.Client
	key string
}

func (h *hash) Push(uuid string, data *dao.Hash) (bool, error) {
	datas, _ := json.Marshal(data)
	ok, err := h.db.HSet(h.key, uuid, datas).Result()
	return ok, err
}

func (h *hash) Pull(uuid string) (data *dao.Hash, err error) {
	json_data, err := h.db.HGet(h.key, uuid).Result()
	if err != nil {
		return
	}
	//fmt.Println(json_data)
	var result dao.Hash
	err = json.Unmarshal([]byte(json_data), &result)
	data = &result
	return
}

func (h *hash) Remove(uuid string) {
	h.db.HDel(h.key, uuid)
}
