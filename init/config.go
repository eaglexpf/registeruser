package init

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"registeruser/entity/global"
)

func init() {
	var configFile string
	flag.StringVar(&configFile, "config_file", "./config.yaml", "your's config file")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
		// 加载依赖
		load()
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	// 加载依赖
	load()
}

func load() {
	// 加载log
	initLog()
	if global.CONFIG.App.Mysql {
		// 加载mysql
		initMysql()
	}
	if global.CONFIG.App.Redis {
		// 加载redis
		initRedis()
	}
}

func Unload() {
	closeMysql()
	closeRedis()
}
