package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	key     []byte        = []byte{}
	expires time.Duration = 2 * time.Hour
)

func init() {
	key = []byte(viper.GetString("jwt.key"))
	i := viper.GetInt("jwt.expires")
	if i > 0 {
		expires = time.Duration(i) * time.Second
	}
}

// 签名体
type Claims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

func (c *Claims) IsExpired() bool {
	return c != nil && time.Now().Unix() > c.ExpiresAt
}

// 生成JWT
func GenerateJWT(uid int) (string, int64) {
	if uid == 0 {
		fmt.Println("生成JWT时，uid为0")
	}
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(key)
	if err != nil {
		panic(err)
	}

	return token, claims.ExpiresAt
}

// 解析JWT，注意过期不是错误，自己判断 IsExpired()
func ParseJWT(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	var claims *Claims
	if tokenClaims != nil {
		claims = tokenClaims.Claims.(*Claims)
	}

	return claims, err
}
