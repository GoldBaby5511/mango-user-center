package middleware

import (
	"fmt"
	"mango-user-center/pkg/response"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		err, ok := recovered.(error)
		if !ok {
			err = fmt.Errorf("%v", recovered)
		}
		response.Echo(c, nil, err)
	})
}
