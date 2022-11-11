package dao

import (
	"fmt"
	"middleground/database"
)

type UserDao struct{}

func (U *UserDao) GetUserList() []database.USER {
	var list []database.USER
	var db database.Init

	s := "select * from user"
	rows, err := db.DB("db").Query(s)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var user database.USER

		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}

		list = append(list, user)
	}

	return list
}
