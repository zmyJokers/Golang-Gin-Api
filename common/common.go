package common

import "time"

/*
	GetCurrentTime 获取当前时间
*/
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
