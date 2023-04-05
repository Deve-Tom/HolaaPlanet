package configs

import (
	"github.com/spf13/viper"
)

// InitRootConfig
// Maintainers:贺胜  Times:2023-04-05
// Part 1:用于读取根目录下的配置文档
// Part 2:主要使用了viper方法加载配置文件，可单独调用该方法，以便在所需地方使用viper提供的相应方法获取文档信息/*
func InitRootConfig() {
	// 加载"./"目录下的"conf.yaml"配置文件
	viper.AddConfigPath("./")
	viper.SetConfigFile("conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Load config file false")
		return
	}
}
