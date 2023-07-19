package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	OpenId  string `json:"openid"`
	BaseUrl string `json:"baseurl"`
}

var configName = ".memos.json"
var openIdKey = "openid"
var baseConfigFile = fmt.Sprintf("%s/%s", GetConfigFilePath(), configName)

// init 初始化配置文件
func init() {
	viper.SetConfigFile(baseConfigFile)
	viper.SetConfigType("json")
	// 读取配置文件
	_ = viper.ReadInConfig()
}

// Set 设置 openId
func Set(key, value string) (string, error) {
	// 写配置文件
	viper.Set(key, value)
	err := viper.WriteConfigAs(baseConfigFile)
	if err != nil {
		return "set openId failed", fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return value, nil
}

// GetConfigFilePath 获取配置文件路径
func GetConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return homeDir
}

// Get 根据key获取值
func Get(key string) string {
	value := viper.GetString(key)
	if value == "" {
		panic(fmt.Errorf("please set memos openId"))
	}
	return value
}

// GetOpenId 获取 openId
func GetOpenId() string {
	value := viper.GetString(openIdKey)
	if value == "" {
		panic(fmt.Errorf("please set memos openId"))
	}
	return value
}

func Reset() (string, error) {
	viper.Reset()
	err := viper.WriteConfigAs(baseConfigFile)
	if err != nil {
		return "reset openId failed", fmt.Errorf("Reset config error config file: %s \n", err)
	}
	return "reset openId success", nil
}
