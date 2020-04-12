// 全局变量
package global

import (
	"database/sql"
	"github.com/go-redis/redis"
)

var (
	// CONFIG 配置文件内容
	CONFIG config
	// DB 全局sql连接池
	DB *sql.DB
	// REDIS 全局redis
	REDIS *redis.Client
)

// 配置文件内容
type config struct {
	App   config_app   `json:"app"`
	Log   config_log   `json:"log"`
	Jwt   config_jwt   `json:"jwt"`
	Mysql config_mysql `json:"mysql"`
	Redis config_redis `json:"redis"`
}

// app配置内容
type config_app struct {
	Addr  string `json:"addr"`
	Mysql bool   `json:"mysql"`
	Redis bool   `json:"redis"`
}

// log配置内容
type config_log struct {
	File config_log_file `json:"file"`
}

// log写入文件的配置内容
type config_log_file struct {
	Status  bool   `json:"status"`
	Path    string `json:"path"`
	MaxSize int    `json:"max_size" mapstructure:"max_size"`
	DayNum  int    `json:"day_num" mapstructure:"day_num"`
	FileNum int    `json:"file_num" mapstructure:"file_num"`
}

// jwt的配置内容
type config_jwt struct {
	Sign    string `json:"sign"`
	Express int    `json:"express"`
	Issuer  string `json:"issuer"`
}

// mysql的配置内容
type config_mysql struct {
	Addr        string `json:"addr"`
	DbName      string `json:"db_name" mapstructure:"db_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Config      string `json:"config"`
	Prefix      string `json:"prefix"`
	MaxIdleConn int    `json:"max_idle_conn" mapstructtrue:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn" mapstructure:"max_open_conn"`
	LogMode     bool   `json:"log_mode"`
}

// redis的配置内容
type config_redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
