package controller

import (
	"github.com/gin-gonic/gin"
	"middleground/app/dao"
)

type TestController struct {
	UserDao dao.UserDao
}

/*
	func LogSendAction
	日志记录
*/
func (T *TestController) TestAction(c *gin.Context) {
	list := T.UserDao.GetUserList()
	res := Response{
		Code: 200,
		Msg:  "success",
		Data: list,
	}

	res.Success(c)

	return
	//获取json中的key,注意使用 . 访问
	//name := params.Name
	//age := params.Age
	//
	//fmt.Println(name, age)
}
