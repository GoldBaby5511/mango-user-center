package model

import (
	"time"
)

var (
	UserLoginPtr *UserLogin
)

type UserLogin struct {
	BaseID
	UserId          int      `json:"user_id" gorm:"index:idx_userid; not null; default:0; comment:用户ID;"`
	RefreshToken    string   `json:"refresh_token" gorm:"size:32; not null; default:''; comment:刷新token;"`
	RefreshTime     Datetime `json:"refresh_time" gorm:"type:timestamp; comment:刷新时间;"`
	DeviceOS        uint8    `json:"device_os" gorm:"type:tinyint unsigned; not null; default:0; comment:设备系统 0未知1苹果2安卓3Web;"`
	DeviceOSVersion string   `json:"device_os_version" gorm:"size:64; not null; default:''; comment:设备系统版本号;"`
	DeviceID        string   `json:"device_id" gorm:"size:64; not null; default:''; comment:设备ID;"`
	Version         string   `json:"version" gorm:"size:64; not null; default:''; comment:版本号;"`
	IP              int      `json:"ip" gorm:"type:int; not null; default:0; comment:IP地址;"`
	// City            string   `json:"city" gorm:"size:16; not null; default:''; comment:城市;"`
	BaseTimes
}

func (*UserLogin) GetByRecord(userId int, refreshToken string) UserLogin {
	var row UserLogin
	DebugDB().Model(UserLoginPtr).Where("user_id = ? and refresh_token = ?", userId, refreshToken).Last(&row)
	return row
}

func (r *UserLogin) IsExpired() bool {
	return time.Since(r.RefreshTime.ToTime()).Hours() > 7*24
}
