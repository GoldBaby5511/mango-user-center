package model

import (
	"mango-user-center/config"
	"mango-user-center/pkg/db"

	"gorm.io/gorm"
)

func DB() *gorm.DB {
	if config.IsDebug {
		return db.GetDB("default").Debug()
	}
	return db.GetDB("default")
}

func DebugDB() *gorm.DB {
	return db.GetDB("default").Debug()
}

func ForestPropertyDB() *gorm.DB {
	return db.GetDB("forest_property")
}
