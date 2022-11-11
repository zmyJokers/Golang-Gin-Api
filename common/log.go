package common

import (
	"encoding/json"
	"middleground/app/service"
)

type Log struct {
	Time     string `json:"time"`
	Project  string `json:"project"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Request  any    `json:"request"`
	Response any    `json:"response"`
}

type ErrorLog struct {
	Time     string `json:"time"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Line     int    `json:"line"`
	Request  any    `json:"request"`
	Response any    `json:"response"`
}

func Success(L *Log) {
	data, _ := json.Marshal(*L)
	_ = service.MQSend("db", "log_middleground", "log_exchange_request", "log_middleground_route_request", data)
}

func Error(E *ErrorLog) {
	data, _ := json.Marshal(*E)
	_ = service.MQSend("db", "log_middleground_error", "log_exchange_error", "log_middleground_route_error", data)
}

func (l *Log) Info() {

}
