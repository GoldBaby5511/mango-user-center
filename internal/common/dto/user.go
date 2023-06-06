package dto

import (
	"mango-user-center/internal/model"
	"mango-user-center/pkg/token"
)

type UserInfo struct {
	SysId   int    `json:"-"`
	Account string `json:"account"`
	// Nickname     string `json:"nickname"`
	model.Avatar  `json:"avatar"`
	WalletAddress string `json:"wallet_address"`
}

type UserSendCodeReq struct {
	Type    uint8  `json:"type" form:"type" validate:"required"`
	Account string `json:"account" form:"account" validate:"required,email"`
	Ip      string
}

type UserSendCodeResp = Null

type UserCheckCodeReq struct {
	// Type       uint8  `json:"type" form:"type" validate:"required"`
	Account    string `json:"account" form:"account" validate:"required,email"`
	VerifyCode string `json:"verify_code" form:"verify_code" validate:"required"`
}

type UserCheckCodeResp = Null

type UserSignupReq struct {
	Account    string `json:"account" form:"account" validate:"required,email,gte=8,lte=32"`
	Password   string `json:"password" form:"password" validate:"required"`
	VerifyCode string `json:"verify_code" form:"verify_code" validate:"required"`

	WalletCategory uint8  `json:"wallet_category" form:"wallet_category" validate:"omitempty,max=4"`
	WalletAddress  string `json:"wallet_address" form:"wallet_address" `
	// Nickname         string `json:"nickname" form:"nickname"`
	// SiteId           uint8  `json:"site_id" form:"site_id"`
	// ChannelId        uint8  `json:"channel_id" form:"channel_id"`

	InviterId string `json:"inviter_id" form:"inviter_id"`
}

type UserSignupResp = UserLoginByWalletResp

type UserVerifyReq struct {
	Account    string `json:"account" form:"account" validate:"required,email,lte=32"`
	VerifyCode string `json:"verify_code" form:"verify_code" validate:"required"`
}

type UserVerifyResp = token.Token

type UserLoginReq struct {
	Account  string `json:"account" form:"account" validate:"required,email,gte=8,lte=32"`
	Password string `json:"password" form:"password" validate:"required"`

	InviteCode string `json:"invite_code" form:"invite_code"`

	DeviceOS        uint8  `json:"device_os" form:"device_os"`
	DeviceOSVersion string `json:"device_os_version" form:"device_os_version"`
	DeviceID        string `json:"device_id" form:"device_id"`
	Version         string `json:"version" form:"version"`
	IP              string
}

type UserLoginResp struct {
	Token token.Token `json:"token"`
	User  UserInfo    `json:"user"`
}

type UserRefreshTokenReq struct {
	AccessToken  string `json:"access_token" form:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}

type UserRefreshTokenResp = token.Token

type UserChangePasswordReq struct {
	// AccessToken string `json:"access_token" form:"access_token" `
	Account     string `json:"account" form:"account" validate:"required"`
	VerifyCode  string `json:"verify_code" form:"verify_code" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}

type UserPostInfoReq struct {
	Nickname *string `json:"nickname" form:"nickname" validate:"omitempty,lte=32"`
	Avatar   *string `json:"avatar" form:"avatar" validate:"omitempty,lte=255"`
	Sex      *uint8  `json:"sex" form:"sex" validate:"omitempty,min=1,max=2"`
	Age      *uint8  `json:"age" form:"age" validate:"omitempty,min=1,max=130"`
}

type UserPostInfoResp = UserInfo

type UserGetInfoReq = Null

type UserGetInfoResp struct {
	model.BaseID
	UserId        int    `json:"user_id"`
	GameId        int    `json:"game_id"`
	Account       string `json:"account"`
	Nickname      string `json:"nickname"`
	model.Avatar  `json:"avatar"`
	WalletAddress string `json:"wallet_address"`
	PrivateKey    string `json:"private_key"`
}

type UserBindWalletReq struct {
	Category uint8  `json:"category" form:"category" validate:"omitempty,max=4"`
	Address  string `json:"address" form:"address" `
}

type UserBindWalletResp = Null

type UserGetWalletNonceReq struct {
	Category uint8  `json:"category" form:"category" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UserGetWalletNonceResp struct {
	Nonce string `json:"nonce"`
}

type UserLoginByEmailReq struct {
	Category  uint8  `json:"category" form:"category" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Signature string `json:"signature" form:"signature" validate:"required"`
}

type UserLoginByWalletReq struct {
	Category  uint8  `json:"category" form:"category" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Signature string `json:"signature" form:"signature" validate:"required"`
}

type UserLoginByWalletResp struct {
	Token token.Token `json:"token"`
	User  UserInfo    `json:"user"`
}

type UserOptionResp struct {
	Avatars []model.Avatar `json:"avatars"`
}

type UserSendInviteCodeReq struct {
	Accounts string `json:"accounts" form:"accounts" validate:"required"`
}

type GetSelfInvitationInfoResp struct {
	Completed int `json:"completed"`
	Goals     int `json:"goals"`
}
