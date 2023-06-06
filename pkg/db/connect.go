package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	writer io.Writer = os.Stdout
	// 慢查询阈值
	slowThresholdMs time.Duration = 300 * time.Millisecond
	// 连接参数
	maxIdleConns = 10
	maxOpenConns = 0
	maxLifeTime  = time.Hour
)

func Connect() {
	for k := range dbConfigs {
		switch dbConfigs[k].Driver {
		case "mysql":
			dbs[k] = connectMySQL(dbConfigs[k])
		// case "postgres":
		// 	dbs[k] = connectPostgres(dbConfigs[k])
		default:
			panic("unsupported driver: " + dbConfigs[k].Driver)
		}
	}
	// SetLogWriter(os.Stdout)
}

func SetLogWriter(w io.Writer) {
	writer = w
}

func connectMySQL(conf config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=False&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		fmt.Println("DB dsn : " + dsn)
		panic("failed to connect database: " + err.Error())
	}

	setConnParams(db)

	setLogger(db)

	return db
}

// func connectPostgres(conf config) *gorm.DB {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s TimeZone=%s",
// 		conf.Host,
// 		conf.Port,
// 		conf.User,
// 		conf.DBName,
// 		conf.Password,
// 		time.Local.String(),
// 	)

// 	config := &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix:   conf.TablePrefix,
// 			SingularTable: true,
// 		},
// 		DisableForeignKeyConstraintWhenMigrating: true,
// 	}

// 	db, err := gorm.Open(postgres.Open(dsn), config)
// 	if err != nil {
// 		panic("failed to connect database: " + err.Error())
// 	}

// 	setConnParams(db)

// 	return db
// }

func setConnParams(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(maxLifeTime)
}

func setLogger(db *gorm.DB) {
	db.Logger = logger.New(
		log.New(writer, "\r\n", log.Ltime),
		logger.Config{
			SlowThreshold:             slowThresholdMs,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}
