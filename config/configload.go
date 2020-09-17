package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Database `json:"database" yaml:"database"`
	Redis    `json:"redis"  yaml:"redis"`
	Log      `json:"log"  yaml:"log"`
}

type Database struct {
	DbConnection  string `yaml:"DbConnection"`
	DbHost        string
	DbPort        int
	DbDatabase    string
	DbUsername    string
	DbPassword    string
	MaxConnection int
	MaxIdle       int
}

type Redis struct {
	RedisHost     string
	RedisPassword string
	RedisPort     int
}

type Log struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

//数据库，redis相关配置
var InitConfig GlobalConfig

func init() {
	//设置文件名
	v := viper.New()
	v.SetConfigName("dev_env")
	//设置后缀类型
	v.SetConfigType("yaml")
	//设置路径
	v.AddConfigPath("env")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fata error config file :%s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&InitConfig); err != nil {
		fmt.Printf("err:%s", err)
	}
}
