package db

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	dbs       = make(map[string]*gorm.DB)
	dbConfigs map[string]config
)

type config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string

	TablePrefix string
}

func init() {
	viper.UnmarshalKey("dbs", &dbConfigs)
}

func GetDB(key string) *gorm.DB {
	return dbs[key]
}
