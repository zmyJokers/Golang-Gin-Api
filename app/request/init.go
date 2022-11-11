package request

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"middleground/app/middleware"
)

func Init(c *gin.Context, data any) {
	err := c.BindJSON(data)
	if err != nil {
	}

	middleground.Params, _ = json.Marshal(data)
	return
}
