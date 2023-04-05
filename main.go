package main

import (
	"HolaaPlanet/configs"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	configs.InitRootConfig()
	fmt.Printf(viper.GetString("mysql.user"))
}
