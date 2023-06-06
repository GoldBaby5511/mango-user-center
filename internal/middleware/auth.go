package middleware

import (
	"mango-user-center/pkg/response"
	"mango-user-center/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

// 使用方式：Header中增加 Authorization: Bearer <token>
const authorization = "Authorization"

func Auth(checkExpires bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader(authorization)
		index := strings.IndexByte(auth, ' ')
		if auth == "" || index < 0 {
			response.Echo(c, nil, response.Msg("Authorization is required."))
			return
		}

		accessToken := auth[index+1:]

		claims, err := token.ParseJWT(accessToken)
		if err != nil && !checkExpires && strings.HasPrefix(err.Error(), "token is expired") {
			err = nil
		}
		if err != nil {
			response.Echo(c, nil, response.Msg(err.Error()))
			return
		}

		// 上下文带上用户ID
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
