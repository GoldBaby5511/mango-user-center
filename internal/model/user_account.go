package model

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

	Password string `json:"-" gorm:"column:passwd; size:64; not null; default:''; comment:密码;"`
	Salt     string `json:"-" gorm:"size:10; not null; default:''; comment:密码盐;"`

	BaseTimes
}

func (*UserAccount) GetByAccount(account string) (row UserAccount) {
	DB().Where("account = ?", account).Take(&row)
	return
}
