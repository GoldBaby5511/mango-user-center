package middleware

import (
	"errors"
	"strings"
	"sync"
	"mango-user-center/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var (
	limiter sync.Map
)

// 限流
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// 非内网请求限流
		if !isIntranetIP(ip) {
			l, ok := limiter.Load(ip)
			if ok {
				if !l.(*rate.Limiter).Allow() {
					logrus.WithField("ip", ip).Warn("超过限流")
					// response.Echo(c, nil, errors.New("Too Many Requests"))
					// return
				}
			} else {
				// 桶最大容量20，每秒产生令牌数5
				x := rate.NewLimiter(20, 5)
				limiter.Store(ip, x)
			}
		}

		c.Next()
	}
}

// 内网访问
func MustIntranet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !isIntranetIP(ip) {
			response.Echo(c, nil, errors.New("非内网访问: "+c.Request.URL.Path))
			return
		}

		c.Next()
	}
}

func isIntranetIP(ip string) bool {
	return ip == "127.0.0.1" || strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "172.") || strings.HasPrefix(ip, "10.")
}
