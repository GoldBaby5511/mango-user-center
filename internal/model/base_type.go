package model

import (
	"mango-user-center/config"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// mysql的int unsigned是4个字节，对应go的uint32。  big int unsigned 对应uint64。 用int是因为方便
type BaseID struct {
	SysId int `json:"sys_id" gorm:"primaryKey; <-:false; type:int unsigned auto_increment; comment:主键;"`
}

// 替代 xx.Id > 0
func (b BaseID) IsValid() bool {
	return b.SysId > 0
}

// 公共字段，如果不需要建立索引，直接使用它
// 禁用了gorm更改这两个字段，数据库自动维护
type BaseTimes struct {
	CreateTime Datetime `json:"create_time" gorm:"<-:false; type:datetime; not null; default:now(); comment:创建时间;"`
	ChangeTime Datetime `json:"change_time" gorm:"<-:false; type:datetime; not null; default:now() ON UPDATE now(); comment:更新时间;"`
}

// 使用Datetime，须关闭gorm连接时的Parsetime，且数据库该字段自动为当前时间戳。
type Datetime string

func (d Datetime) ToTime() time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", string(d), time.Local)
	if err != nil {
		logrus.Warn("parse time error: ", err)
	}
	return t
}

func (d Datetime) Now() Datetime {
	return Datetime(time.Now().Format("2006-01-02 15:04:05"))
}

// 自动填充URL
type URL string

func (u *URL) Pad() {
	*u = URL(config.App.Host + string(*u))
}

func (u *URL) Trim() {
	*u = URL(strings.TrimPrefix(string(*u), config.App.Host))
}

func (u *URL) AfterFind(tx *gorm.DB) error {
	u.Pad()
	return nil
}

func (u *URL) BeforeUpdate(tx *gorm.DB) error {
	u.Trim()
	return nil
}

type Avatar = URL
