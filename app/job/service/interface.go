package service

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"registeruser/app/job/entity/dao"
	"registeruser/app/job/entity/request"
	"registeruser/app/job/model"
	"registeruser/app/job/model/redis"
	"strings"
	"time"
)

type JobService interface {
	CreateJob(ctx context.Context, data *request.RequestRegisterJob) (ok bool, err error)
	ExpireTask()
	RunTask()
	SubTask()
	SetList()
}

func NewJobService() JobService {
	return &task{
		hashModel:    redis.NewHash(),
		sortSetModel: redis.NewSortSet(),
		listModel:    redis.NewList(),
		setModel:     redis.NewSet(),
		chanelModel:  redis.NewChanel(),
	}
}

type task struct {
	hashModel    model.Hash
	sortSetModel model.SortSet
	listModel    model.List
	setModel     model.Set
	chanelModel  model.Chanel
}

func (t *task) CreateJob(ctx context.Context, data *request.RequestRegisterJob) (ok bool, err error) {
	run_time := time.Now().Unix() + data.Delay
	uuid_str := fmt.Sprintf("%s", uuid.NewV4())
	job := &dao.Hash{
		UUID:        uuid_str,
		URI:         data.URI,
		Method:      data.Method,
		Data:        data.Data,
		Time:        run_time,
		LastTime:    run_time,
		Delay:       data.Delay,
		Bomb:        data.Bomb,
		Num:         0,
		SuccessData: data.SuccessData,
	}
	// 将任务压入hash中
	ok, err = t.hashModel.Push(uuid_str, job)
	if err != nil || !ok {
		return
	}
	// 将任务id压入有序集合中
	err = t.sortSetModel.Push(uuid_str, time.Now().Unix()+data.Delay)
	return
}

// 检测有序集合中是否有可执行任务
func (t *task) ExpireTask() {
	// 获取可以执行的任务的有序集合
	data, err := t.sortSetModel.Pull()
	if err != nil {
		return
	}
	for _, v := range data {
		// 判断任务id在集合中是否存在
		// 存在则忽略；不存在则压入集合与列表中
		//fmt.Println(t.setModel.Exist(v))
		if is, err := t.setModel.Exist(v); err == nil && !is {
			// 将任务id压入唯一集合
			err = t.setModel.Push(v)
			if err != nil {
				continue
			}
			// 将任务id压入列表
			err := t.listModel.Push(v)
			if err != nil {
				continue
			}
		}

	}
	// 发送订阅消息
	//t.chanelModel.Pub("success")
}

// 执行任务
func (t *task) RunTask() {
	for {
		fmt.Println("run task")
		// 从列表中读取任务id
		list_uuid, err := t.listModel.Pull() //没数据时阻塞
		if err != nil {
			fmt.Println("list数据读取失败", list_uuid, err)
			break
		}
		// 在唯一集合中删除并判断任务id是否存在
		is, err := t.setModel.Remove(list_uuid)
		if err != nil || !is {
			// 任务不在唯一集合中或者删除失败，直接丢弃
			// 丢弃后timer服务会将其再次放入list与集合内
			// continue执行下一个任务
			continue
		}

		// 从hash中领取相应的任务
		hash, err := t.hashModel.Pull(list_uuid)
		if err != nil {
			// 没有任务；删除任务有序集合内的id
			t.sortSetModel.Remove(list_uuid)
			continue
		}

		// 执行任务
		// 执行http任务
		//fmt.Println(time.Now().String(), hash.URI, hash.Method, hash.Num)
		run_status := false
		for {
			client := &http.Client{}
			req, err := http.NewRequest(strings.ToUpper(hash.Method), hash.URI, bytes.NewReader([]byte(hash.Data)))
			if err != nil {
				fmt.Println("http 请求失败：", err)
				break
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("http 请求失败：", err)
				break
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("http 请求失败：", err)
				break
			}

			if hash.SuccessData == "" {
				if string(body) != "success" {
					break
				}
			} else {
				if string(body) != hash.SuccessData {
					break
				}
			}

			run_status = true
			break
		}
		// 任务执行成功后，清理数据；删除hash，有序集合内的数据
		t.sortSetModel.Remove(list_uuid)
		t.hashModel.Remove(list_uuid)

		if hash.Bomb {
			// 定时任务，重复执行
			hash.Num += 1
			run_time := hash.Time + hash.Num*hash.Delay
			now_time := time.Now().Unix()
			if run_time < now_time {
				run_time = now_time + hash.Delay
			}
			uuid_str := fmt.Sprintf("%s", uuid.NewV4())
			hash.UUID = uuid_str
			hash.LastTime = run_time
			_, _ = t.hashModel.Push(uuid_str, hash)
			_ = t.sortSetModel.Push(uuid_str, run_time)
			// 如果是定时任务，不进行失败后的重试
			break
		}
		if !run_status {
			hash.Num += 1
			var num_time int64 = 1
			var i int64
			for i = 0; i < hash.Num; i++ {
				num_time *= 2
			}
			if num_time > 24*3600 {
				// 最多重试时间：24小时
				break
			}
			run_time := hash.Time + num_time
			now_time := time.Now().Unix()
			if run_time < now_time {
				// 应该执行的时间小于当前时间
				run_time = now_time + num_time
			}
			uuid_str := fmt.Sprintf("%s", uuid.NewV4())
			hash.UUID = uuid_str
			hash.LastTime = run_time
			_, _ = t.hashModel.Push(uuid_str, hash)
			_ = t.sortSetModel.Push(uuid_str, run_time)
		}
	}

}

// 将集合中的数据转存到list中
func (t *task) SetList() {
	for {
		set_uuid, err := t.setModel.Pop()
		if err != nil {
			fmt.Println("从集合中取出数据失败： %v", err)
			break
		}
		err = t.listModel.Push(set_uuid)
		if err != nil {
			fmt.Println("向list中压入数据失败： %v", err)
			break
		}
		fmt.Println("list压入数据成功：", set_uuid)
	}

}

// 等待订阅消息，接收到订阅消息后压入list
func (t *task) SubTask() {
	msgChan := make(chan string)
	go t.chanelModel.Sub(msgChan)
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("server", msg)
			// 接收到订阅消息；从集合中取出数据压入list中
			t.SetList()
		}
	}
}
