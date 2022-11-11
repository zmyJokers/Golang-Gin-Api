package log

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"middleground/app/controller"
	"middleground/app/request"
	"middleground/app/request/logRequest"
	"middleground/app/service"
)

type Controller struct {
	LogParams logRequest.LogSendAction
}

/*
	func LogSendAction
	日志记录
*/
func (C *Controller) LogSendAction(c *gin.Context) {
	// 参数初始化
	request.Init(c, &C.LogParams)

	queueName, ok := C.LogParams.RabbitMQ["queue"] // 队列名
	if !ok {
		panic("缺少参数：队列名")
	}

	exchangeName, ok := C.LogParams.RabbitMQ["exchange"] // 交换机
	if !ok {
		panic("缺少参数：交换机")
	}

	routingKey, ok := C.LogParams.RabbitMQ["routing_key"] // 路由
	if !ok {
		panic("缺少参数：路由")
	}

	data, err := json.Marshal(C.LogParams.Response)
	if err != nil {
		panic(err.Error())
	}

	// 推送队列消息
	err = service.MQSend("db", queueName, exchangeName, routingKey, data)
	if err != nil {
		panic(err.Error())
	}

	res := controller.Response{
		Code: 200,
		Msg:  "success",
		Data: "",
	}

	// 输出接口数据
	res.Success(c)
}
