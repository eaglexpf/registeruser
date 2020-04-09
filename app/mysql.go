package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"registeruser/entity/global"
	"registeruser/log"
)

func InitMysql() {
	admin := global.CONFIG.Mysql
	if db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		admin.Username,
		admin.Password,
		admin.Addr,
		admin.DbName,
		admin.Config)); err != nil {
		log.Fatalf("Mysql数据库链接异常:%v", err)
	} else {
		global.DB = db
		global.DB.SingularTable(true)
		global.DB.LogMode(admin.LogMode)
		global.DB.DB().SetMaxIdleConns(admin.MaxIdleConn)
		global.DB.DB().SetMaxOpenConns(admin.MaxOpenConn)
		log.Info("Mysql数据库已链接...")
	}
}
