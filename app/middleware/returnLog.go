package middleground

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"middleground/app/controller"
	"middleground/common"
	"net/http"
	"runtime"
)

var Params []byte

// CustomResponseWriter 重构gin ResponseWriter
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Return() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = res

		c.Next()

		var log = &common.Log{
			Time:     common.GetCurrentTime(),
			Project:  "oms5",
			Path:     c.Request.URL.Path,
			Method:   c.Request.Method,
			Request:  string(Params),
			Response: res.body.String(),
		}
		common.Success(log) // 推送日志
	}
}

/*
	Recovery 异常监控
*/
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			_, fileName, line, _ := runtime.Caller(2)
			if err := recover(); err != nil {
				var log = &common.ErrorLog{
					Time:     common.GetCurrentTime(),
					Path:     fileName,
					Method:   c.Request.Method,
					Line:     line,
					Request:  "",
					Response: err,
				}
				fmt.Println(log)

				common.Error(log) // 推送日志

				// 输出接口数据
				res := controller.Response{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("%s", err),
					Data: "",
				}
				res.Error(c)
			}
		}()

		c.Next()
	}
}
