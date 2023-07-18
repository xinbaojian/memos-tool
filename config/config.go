package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var configName = ".memos.yaml"
var openIdKey = "openid"

// SetOpenId 设置 openId
func SetOpenId(openId string) (string, error) {
	baseConfigFile := fmt.Sprintf("%s/%s", GetConfigFilePath(), configName)
	viper.SetConfigFile(baseConfigFile)
	//viper.SetConfigType("yaml")
	// 读取配置文件
	_ = viper.ReadInConfig()
	// 写配置文件
	viper.Set(openIdKey, openId)
	err := viper.WriteConfigAs(baseConfigFile)
	if err != nil {
		return "set openId failed", fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return openId, nil
}

// GetConfigFilePath 获取配置文件路径
func GetConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return homeDir
}

// GetOpenId 获取 openId
func GetOpenId() string {
	baseConfigFile := fmt.Sprintf("%s/%s", GetConfigFilePath(), configName)
	viper.SetConfigFile(baseConfigFile)
	_ = viper.ReadInConfig()
	openId := viper.GetString(openIdKey)
	if openId == "" {
		panic(fmt.Errorf("please set memos openId"))
	}
	return openId
}

func Reset() (string, error) {
	baseConfigFile := fmt.Sprintf("%s/%s", GetConfigFilePath(), configName)
	viper.SetConfigFile(baseConfigFile)
	viper.Reset()
	err := viper.WriteConfigAs(baseConfigFile)
	if err != nil {
		return "set openId failed", fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return "reset openId success", nil
}
