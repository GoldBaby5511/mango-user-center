package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	// 启动参数
	Flag struct {
		ConfigFile string
		IsMigrate  bool
	}
	// 配置解析
	App struct {
		Env    string
		Debug  bool
		System string
		Host   string
		Port   string
		LogDir string
	}
	// 本地其他服务调用
	RPC struct {
		CommonLogic string
	}
)

var (
	IsLocal bool
	IsDev   bool
	IsTest  bool
	IsProd  bool

	IsDebug bool
)

func init() {
	flag.StringVar(&Flag.ConfigFile, "c", "./config/config.yaml", "指定配置文件")
	flag.BoolVar(&Flag.IsMigrate, "m", false, "是否数据库迁移")
	flag.Parse()

	// 读取配置
	viper.SetConfigFile(Flag.ConfigFile)
	ext := Flag.ConfigFile[strings.LastIndexByte(Flag.ConfigFile, '.')+1:]
	viper.SetConfigType(ext)

	LoadConfig()
}

func LoadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 解析配置
	viper.UnmarshalKey("app", &App)
	viper.UnmarshalKey("rpc", &RPC)

	// 环境变量
	IsLocal = App.Env == "local"
	IsDev = App.Env == "dev"
	IsTest = App.Env == "test"
	IsProd = App.Env == "prod"
	IsDebug = App.Debug

	fmt.Printf("当前环境：%s，服务端口：%s \n", App.Env, App.Port)
}
