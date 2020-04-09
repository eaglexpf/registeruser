package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	CONFIG config
	DB     *gorm.DB
	REDIS  *redis.Client
)

type config struct {
	App   config_app   `json:"app"`
	Log   config_log   `json:"log"`
	Mysql config_mysql `json:"mysql"`
	Redis config_redis `json:"redis"`
}

type config_app struct {
	Addr  string `json:"addr"`
	Mysql bool   `json:"mysql"`
	Redis bool   `json:"redis"`
}

type config_log struct {
	File config_log_file `json:"file13"`
}

type config_log_file struct {
	Status  bool   `json:"status"`
	Path    string `json:"path"`
	MaxSize int    `json:"max_size"`
	DayNum  int    `json:"day_num"`
	FileNum int    `json:"file_num"`
}

type config_mysql struct {
	Addr        string `json:"addr"`
	DbName      string `json:"db_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Config      string `json:"config"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
	LogMode     bool   `json:"log_mode"`
}

type config_redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
