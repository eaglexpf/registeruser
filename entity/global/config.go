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
	App   App   `json:"app"`
	Log   Log   `json:"log"`
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type App struct {
	Addr  string `json:"addr"`
	Mysql bool   `json:"mysql"`
	Redis bool   `json:"redis"`
}

type Log struct {
	Path string `json:"path"`
}

type Mysql struct {
	Addr        string `json:"addr"`
	DbName      string `json:"db_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Config      string `json:"config"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
	LogMode     bool   `json:"log_mode"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
