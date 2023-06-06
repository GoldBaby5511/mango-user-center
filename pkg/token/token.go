package token

import (
	"math/rand"
	"time"
)

const (
	charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	charsetLen            = len(charset)
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Token struct {
	// 有效期较长。后端需要存储，自己限制其过期时间。
	RefreshToken string `json:"refresh_token"`
	// 有效期短。是个JWT， 后端不需要存储
	AccessToken string `json:"access_token"`
	// JWT的过期时间，前端存储，判断是否需要刷新
	ExpireAt int64 `json:"expire_at"`
}

func GenerateTokens(uid int) Token {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(charsetLen)]
	}
	var token Token
	token.RefreshToken = string(b)
	token.AccessToken, token.ExpireAt = GenerateJWT(uid)
	return token
}
