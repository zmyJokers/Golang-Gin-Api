package database

import (
	"database/sql"
	"fmt"
	"middleground/conf"
	"time"
)

// Connect 数据库连接
var Connect map[string]*sql.DB

type Init struct{}

func (*Init) Mysql() {
	// 读取数据库配置
	mysqlConfig := conf.MysqlConf

	for k, v := range mysqlConfig {
		dsn := v.Username + ":" + v.Password + "@tcp(" + v.Host + ":" + v.Port + ")/" + v.Database + "?charset=utf8&parseTime=True&loc=Local"
		mysqlConn, err := sql.Open("mysql", dsn)
		if err != nil {
			panic("mysql connect error!")
		}

		if mysqlConn.Ping() != nil {
			fmt.Println("mysql ping error!")
		}

		// 最大连接时长
		mysqlConn.SetConnMaxLifetime(time.Minute * 60)
		// 最大连接数
		mysqlConn.SetMaxOpenConns(2000)
		// 空闲连接数
		mysqlConn.SetMaxIdleConns(1000)

		Connect = make(map[string]*sql.DB)
		Connect[k] = mysqlConn
	}
}

func (i *Init) DB(name string) *sql.DB {
	// 检测数据库连接是否正常
	if Connect[name].Ping() != nil {
		i.Mysql()
		i.DB(name)

		if Connect[name].Ping() != nil {
			panic("mysql try two connect error!")
		}
	}

	return Connect[name]
}
