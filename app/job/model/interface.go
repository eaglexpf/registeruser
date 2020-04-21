package model

import (
	"registeruser/app/job/entity/dao"
)

type Hash interface {
	Push(uuid string, data *dao.Hash) (bool, error)
	Pull(uuid string) (data *dao.Hash, err error)
	Remove(uuid string)
}

type SortSet interface {
	Push(uuid string, time_scores int64) (err error)
	Pull() (result []string, err error)
	Remove(uuid string)
}

type List interface {
	Push(uuid string) error
	Pull() (string, error)
}

type Set interface {
	Push(uuid string) error
	Pop() (string, error)
	Remove(uuid string) (is bool, err error)
	Exist(uuid string) (bool, error)
}

type Chanel interface {
	Pub(uuid string)
	Sub(chan<- string)
}
