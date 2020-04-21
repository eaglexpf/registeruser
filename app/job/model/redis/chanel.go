package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"registeruser/app/job/model"
	"registeruser/conf/global"
)

func NewChanel() model.Chanel {
	newChanel := &chanel{
		db:  global.REDIS,
		key: "task_chanel",
	}
	return newChanel
}

type chanel struct {
	db  *redis.Client
	key string
}

func (c *chanel) Pub(uuid string) {
	//err := c.pubClient.Do("PUBLISH", c.key, "success").Err()
	err := c.db.Publish(c.key, uuid).Err()
	fmt.Printf("发布订阅消息123：%s;%v\n", uuid, err)
}

func (c *chanel) Sub(msgChan chan<- string) {
	pubsub := c.db.Subscribe(c.key)
	defer func() {
		fmt.Println(pubsub.Close().Error())
	}()
	_, err := pubsub.Receive()
	if err != nil {
		fmt.Printf("try subscribe channel[test_channel] error[%s]\n",
			err.Error())
		return
	}

	ch := pubsub.Channel()
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		msgChan <- msg.String()
	}
}
