package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"registeruser/conf/global"
	"registeruser/conf/log"
)

// 初始化mysql链接
func initMysql() {
	config := global.CONFIG.Mysql
	if db, err := sql.Open(`mysql`, fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.Username,
		config.Password,
		config.Addr,
		config.DbName,
		config.Config)); err != nil {
		log.Fatalf("Mysql数据库链接异常:%v", err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Fatalf("Mysql数据库ping异常:%v", err)
		}
		global.DB = db
		global.DB.SetMaxIdleConns(config.MaxIdleConn)
		global.DB.SetMaxOpenConns(config.MaxOpenConn)
		log.Info("Mysql数据库已链接...")
	}
}

// 清理mysql链接
func closeMysql() {
	if global.DB == nil {
		return
	}
	err := global.DB.Close()
	if err != nil {
		log.Infof("Mysql数据库关闭连接失败：%v", err)
	}
}
