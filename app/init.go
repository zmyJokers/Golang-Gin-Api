package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"middleground/database"
)

func Init(env string) {
	// 设置运行模式
	if env == "master" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 初始化数据库
	var db database.Init
	db.Mysql()
}
