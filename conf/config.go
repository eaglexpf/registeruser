// 项目配置
package conf

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"registeruser/conf/global"
)

// 加载配置文件，根据配置文件内容加载log，mysql，redis等
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

// 加载依赖；根据配置选择加载log，mysql，redis等
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

// 程序关闭清理，主函数中调用
func Unload() {
	closeMysql()
	closeRedis()
}
