package model

import (
	"time"
)

var (
	UserOnlinePtr *UserOnline
)

type UserOnline struct {
	BaseID
	UserId       int    `json:"user_id" gorm:"type: int unsigned; not null; default 0; comment:用户ID;"`
	OfflineTime  uint32 `json:"offline_time" gorm:"type:int unsigned; not null; default:0; comment:离线时间;"`
	OnlineSecond uint32 `json:"online_second" gorm:"type:int unsigned; not null; default:0; comment:在线时长;"`
	BaseTimes
}

func (*UserOnline) Login(userId int) {
	uo := UserOnline{
		UserId:       userId,
		OfflineTime:  0,
		OnlineSecond: 0,
	}

	DebugDB().Create(&uo)
}

func (*UserOnline) Logout(userId int) {
	var uo UserOnline
	hoursAgo := time.Now().Add(-time.Hour * 10).Format("2006-01-02 15:04:05")
	DebugDB().Where("user_id = ? and create_time >= ?", userId, hoursAgo).Last(&uo)
	DebugDB().Model(&uo).Updates(map[string]interface{}{
		"offline_time":  time.Now().Unix(),
		"online_second": int(time.Since(uo.CreateTime.ToTime()).Seconds()),
	})
}
