package model

import "gorm.io/gorm"

const (
	InvitedStatus = iota + 1
	AttainmentStatus
)

var (
	UserAccountPtr *UserAccount
)

type UserAccount struct {
	BaseID
	// 要求有这个user_id字段
	UserId   int    `json:"user_id" gorm:"uniqueIndex:uk_userid; type:int unsigned; not null; default 0; comment:用户ID;"`
	GameId   int    `json:"game_id" gorm:"uniqueIndex:uk_gameid; type:int unsigned; not null; default:0; comment:游戏ID;"`
	Account  string `json:"account" gorm:"uniqueIndex:uk_account; size:32; not null; default:''; comment:账号;"`
	Nickname string `json:"nickname" gorm:"size:32; not null; default:''; comment:昵称;"`
	State    uint8  `json:"state" gorm:"type:tinyint unsigned; not null; default:3; comment:状态 1正常2冻结3未验证;"`

	Sex       uint8  `json:"sex" gorm:"type:tinyint unsigned; not null; default:3; comment:性别 1男2女3未知;"`
	Age       uint8  `json:"age" gorm:"type:tinyint unsigned; not null; default:0; comment:年龄;"`
	ChannelId uint8  `json:"channel_id" gorm:"type:tinyint unsigned; not null; default:0; comment:注册主渠道;"`
	SiteId    uint8  `json:"site_id" gorm:"type:tinyint unsigned; not null; default:0; comment:注册子渠道;"`
	Balance   uint64 `json:"balance" gorm:"type:bigint unsigned; not null; default:0; comment:余额;"`
	// 注意Avatar是嵌入进来的，所有Avatar都是用这个
	Avatar `json:"avatar" gorm:"type:varchar(255); not null; default:''; comment:头像;"`

	// 钱包
	WalletAddress string `json:"wallet_address" gorm:"index:idx_wallet_address; size:42; not null; default:''; comment:BB钱包地址;"`
	PublicKey     string `json:"public_key" gorm:"size:128; not null; default:''; comment:BB公钥;"`
	PrivateKey    string `json:"-" gorm:"size:64; not null; default:''; comment:BB私钥;"`
	Mnemonic      string `json:"-" gorm:"size:128; not null; default:''; comment:BB助记词;"`

	Password string `json:"-" gorm:"column:passwd; size:64; not null; default:''; comment:密码;"`
	Salt     string `json:"-" gorm:"size:10; not null; default:''; comment:密码盐;"`

	InviterId        int `json:"inviter_id" gorm:"type:int unsigned; not null; default 0;COMMENT '邀请人id';"`
	InvitationStatus int `json:"invitation_status" gorm:"type:tinyint unsigned; not null; default 0; COMMENT '邀请状态：0 无，1 被邀请，2 达到标准';"`

	BaseTimes
}

func (*UserAccount) GetByAccount(account string) (row UserAccount) {
	DB().Where("account = ?", account).Take(&row)
	return
}

func (*UserAccount) GetByInvitationStatus(status int) (rows []*UserAccount) {
	DB().Where("invitation_status = ?", status).Find(&rows)
	return
}

func (u *UserAccount) SetInvitationStatusByUserIds(db *gorm.DB, userIds []int) (err error) {
	err = db.Model(&UserAccount{}).Where("user_id IN ?", userIds).Updates(&UserAccount{InvitationStatus: 2}).Error
	return
}

func (*UserAccount) GetCountByInvitationStatusAndInviterId(status int, inviterId int) (count int64, err error) {
	err = DB().Model(&UserAccount{}).Where("invitation_status = ? and inviter_id = ?", status, inviterId).Count(&count).Error
	return
}

func (*UserAccount) GetRowsByInvitationStatusAndInviterId(status int, inviters []int) (rows []*UserAccount, err error) {
	err = DB().Where("invitation_status = ? and inviter_id IN ?", status, inviters).Find(&rows).Error
	return
}
