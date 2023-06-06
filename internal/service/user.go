package service

import (
	"io/ioutil"
	"mango-user-center/internal/common/dto"
	"mango-user-center/internal/common/enum"
	"mango-user-center/internal/model"
	"mango-user-center/pkg/response"
	"mango-user-center/pkg/token"
	"mango-user-center/pkg/util"
	"strconv"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	signupMu sync.Mutex
)

type User struct{}

// 注册
func (u User) Signup(req dto.UserSignupReq, resp *dto.UserSignupResp) error {
	signupMu.Lock()
	defer signupMu.Unlock()

	// 检测账号重复
	var count int64
	model.DebugDB().Model(model.UserAccountPtr).Where("account = ?", req.Account).Count(&count)
	if count > 0 {
		return response.Msg("Account already exists.")
	}

	err := model.DebugDB().Transaction(func(tx *gorm.DB) error {
		salt := u.newSalt()
		password := util.Encry.HmacSha256(req.Password, salt)
		ua := model.UserAccount{
			State:    enum.UserStateOk,
			Account:  req.Account,
			Avatar:   "/avatar/1.png",
			Password: password,
			Salt:     salt,
		}
		if req.InviterId != "" {
			inviter, err := strconv.Atoi(req.InviterId)
			if err != nil {
				log.WithField("account", req.Account).WithField("req.InviterId", req.InviterId).Error(err)
			} else {
				ua.InviterId = inviter
				ua.InvitationStatus = 1
			}
		}
		err := tx.Create(&ua).Error
		if err != nil {
			return err
		}

		// 分配游戏ID
		err = tx.Model(&ua).Updates(map[string]interface{}{
			"user_id": ua.SysId,
			"game_id": game.DistributeGameId(ua.SysId),
		}).Error
		if err != nil {
			return err
		}

		log.Info("用户注册 : " + req.Account)
		copier.Copy(&resp.User, &ua)
		resp.Token = token.GenerateTokens(ua.SysId)
		return nil
	})

	return err
}

// 登录
func (u User) Login(req dto.UserLoginReq, resp *dto.UserLoginResp) error {
	ua := model.UserAccountPtr.GetByAccount(req.Account)
	if !ua.IsValid() {
		return response.Msg("Account does not exist.")
	}

	// 检查密码
	if ua.Password != util.Encry.HmacSha256(req.Password, ua.Salt) {
		return response.Msg("Please provide correct credentials.")
	}

	switch ua.State {
	// case enum.UserStateUncheck:
	// 	return response.Msg("The account is not verified")
	case enum.UserStateFrozen:
		return response.Msg("The account has been frozen.")
	}

	copier.Copy(&resp.User, &ua)
	resp.Token = token.GenerateTokens(resp.User.SysId)

	go func() {
		ul := model.UserLogin{
			UserId:          ua.UserId,
			RefreshToken:    resp.Token.RefreshToken,
			RefreshTime:     model.Datetime(time.Now().Format("2006-01-02 15:04:05")),
			DeviceOS:        req.DeviceOS,
			DeviceOSVersion: req.DeviceOSVersion,
			DeviceID:        req.DeviceID,
			Version:         req.Version,
			IP:              util.Ip2Int(req.IP),
		}
		err := model.DebugDB().Create(&ul).Error
		if err != nil {
			log.Error(err)
		}
		model.UserOnlinePtr.Login(ua.UserId)
		log.WithField("ip", req.IP).Info("用户登录 : ", ua.UserId)
	}()

	return nil
}

// 刷新token
func (u User) RefreshToken(req dto.UserRefreshTokenReq, resp *dto.UserRefreshTokenResp) error {
	// 从access_token 中获取身份
	claims, err := token.ParseJWT(req.AccessToken)
	if err != nil && !claims.IsExpired() {
		return enum.AccessTokenError
	}

	// 校验
	ul := model.UserLoginPtr.GetByRecord(claims.Uid, req.RefreshToken)
	if !ul.IsValid() {
		return enum.RefreshTokenError
	}
	if ul.IsExpired() {
		return enum.RefreshTokenExpired
	}

	// 生成新的
	*resp = token.GenerateTokens(claims.Uid)
	model.DB().Model(&ul).Updates(&model.UserLogin{
		RefreshToken: resp.RefreshToken,
		RefreshTime:  model.Datetime(time.Now().Format("2006-01-02 15:04:05")),
	})

	go model.UserOnlinePtr.Login(claims.Uid)

	log.Debug("刷新token : ", claims.Uid)

	return nil
}

func (u User) ChangePassword(req dto.UserChangePasswordReq) error {
	ua := model.UserAccountPtr.GetByAccount(req.Account)
	if !ua.IsValid() {
		return response.Msg("Account does not exist.")
	}

	var update model.UserAccount
	update.Salt = u.newSalt()
	update.Password = util.Encry.HmacSha256(req.NewPassword, update.Salt)

	return model.DebugDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ua).Updates(&update).Error; err != nil {
			return err
		}

		return nil
	})
}

// 获取信息
func (u User) GetInfo(uid int, req dto.UserGetInfoReq, resp *dto.UserGetInfoResp) error {
	model.DB().Model(model.UserAccountPtr).Take(resp, uid)

	return nil
}

// 更新信息
func (u User) PostInfo(uid int, req dto.UserPostInfoReq, resp *dto.UserPostInfoResp) error {
	if err := model.DB().Model(model.UserAccountPtr).Take(&resp, uid).Error; err != nil {
		return err
	}

	copier.Copy(resp, &req)

	var update model.UserAccount
	copier.Copy(&update, &req)
	update.SysId = resp.SysId
	err := model.DebugDB().Model(&update).Updates(&update).Error

	return err
}

// 登出
func (u User) Logout(uid int) error {
	// 记录离线
	model.UserOnlinePtr.Logout(uid)
	log.Info("用户下线 : ", uid)

	return nil
}

func (User) newSalt() string {
	return util.Rand.String(10)
}

// 这些头像配置在static目录下，通过nginx外部访问
// 注意，go run 获取不到准确目录，可以使用air
func (User) Option(resp *dto.UserOptionResp) error {
	dir := util.Filer.GetRootDir() + "/static/avatar"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		logrus.WithField("Option", dir).Error(err)
		return nil
	}
	resp.Avatars = make([]model.Avatar, len(files))
	for i := range files {
		resp.Avatars[i] = model.Avatar("/avatar/" + files[i].Name())
		resp.Avatars[i].Pad()
	}
	return nil
}
