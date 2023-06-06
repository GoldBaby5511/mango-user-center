package enum

import (
	"errors"
)

var (
	AccessTokenError    = errors.New("AccessToken Wrong")
	AccessTokenExpired  = errors.New("AccessToken Expired")
	RefreshTokenError   = errors.New("RefreshToken Wrong")
	RefreshTokenExpired = errors.New("RefreshToken Expired")
)
