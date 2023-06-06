package controller

import (
	"strconv"
	"time"
	"mango-user-center/pkg/response"

	"github.com/gin-gonic/gin"
)

type test struct{}

func init() {
	Engine.GET("/test", test{}.Index)
	Engine.GET("/test/panic", test{}.Panic)
}

func (t test) Index(c *gin.Context) {
	sleep := c.Query("sleep")
	if sleep != "" {
		i, _ := strconv.Atoi(sleep)
		time.Sleep(time.Duration(i) * time.Second)
	}

	response.Echo(c, "ok", nil)
}

func (t test) Panic(c *gin.Context) {
	panic("测试panic")
}
