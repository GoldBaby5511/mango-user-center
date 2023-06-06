package controller

import (
	"github.com/gin-gonic/gin"
	"mango-user-center/internal/common/dto"
	"mango-user-center/internal/common/helper"
	"mango-user-center/internal/middleware"
	"mango-user-center/internal/service"
	"mango-user-center/pkg/response"
)

func init() {
	var v1 = Engine.Group("user/v1")
	var u user
	// 无需登录

	v1.POST("signup", u.Signup)                  // 注册
	v1.POST("login", u.Login)                    // 登录
	v1.POST("refresh_token", u.RefreshToken)     // 刷新token
	v1.POST("change_password", u.ChangePassword) // 更改密码

	v1.POST("logout", middleware.Auth(false), u.Logout) // 退出登录 (不检查token过期)
	v1.GET("option", u.Option)                          // 获取配置选项

	// 需要登录
	v1a := v1.Use(middleware.Auth(true))
	v1a.GET("info", middleware.MustIntranet(), u.GetInfo) // 获取个人信息 (仅内网访问)
	v1a.POST("info", u.PostInfo)                          // 修改个人信息

}

// 0字节
type user struct {
	Service service.User
}

func (u user) Signup(c *gin.Context) {
	var req dto.UserSignupReq
	if !helper.BindAndTrim(c, &req) {
		return
	}
	var resp dto.UserSignupResp
	err := u.Service.Signup(req, &resp)
	response.Echo(c, resp, err)
}

func (u user) Login(c *gin.Context) {
	var req dto.UserLoginReq
	if !helper.BindAndTrim(c, &req) {
		return
	}
	req.IP = c.ClientIP()
	var resp dto.UserLoginResp
	err := u.Service.Login(req, &resp)
	response.Echo(c, &resp, err)
}

func (u user) RefreshToken(c *gin.Context) {
	var req dto.UserRefreshTokenReq
	if !helper.Bind(c, &req) {
		return
	}
	var resp dto.UserRefreshTokenResp
	err := u.Service.RefreshToken(req, &resp)
	response.Echo(c, &resp, err)
}

func (u user) ChangePassword(c *gin.Context) {
	var req dto.UserChangePasswordReq
	if !helper.Bind(c, &req) {
		return
	}
	err := u.Service.ChangePassword(req)
	response.Echo(c, dto.Null{}, err)
}

func (u user) GetInfo(c *gin.Context) {
	var req dto.UserGetInfoReq
	if !helper.Bind(c, &req) {
		return
	}
	var resp dto.UserGetInfoResp
	err := u.Service.GetInfo(helper.GetUid(c), req, &resp)
	response.Echo(c, &resp, err)
}

func (u user) PostInfo(c *gin.Context) {
	var req dto.UserPostInfoReq
	if !helper.BindAndTrim(c, &req) {
		return
	}
	var resp dto.UserPostInfoResp
	err := u.Service.PostInfo(helper.GetUid(c), req, &resp)
	response.Echo(c, &resp, err)
}

func (u user) Logout(c *gin.Context) {
	err := u.Service.Logout(helper.GetUid(c))
	response.Echo(c, nil, err)
}

func (u user) Option(c *gin.Context) {
	var resp dto.UserOptionResp
	err := u.Service.Option(&resp)
	response.Echo(c, &resp, err)
}
