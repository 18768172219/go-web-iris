package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/golog"
	"main/common"
)

var G_DB *gorm.DB

/**
 Init gorm 配置初始化 数据库连接
 */
func InitGorm()  {
	var (
		err error
		db *gorm.DB
	    config = common.G_AppConfig
	    url = config.DBPostgres.DBUrl()
	)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		goto ERR
	}
	if err = db.DB().Ping(); err != nil {
		goto ERR
	}
	G_DB = db
	return
ERR:
	golog.Fatalf("postgres的gorm初始化错误", err.Error())
}
