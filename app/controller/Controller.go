package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 返回成功
func (res *Response) Success(c *gin.Context) {
	c.JSON(http.StatusOK, res)
}

// Error 返回失败
func (res *Response) Error(c *gin.Context) {
	c.JSON(http.StatusBadRequest, res)
}
