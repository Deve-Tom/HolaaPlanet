package main

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/routers"
	"fmt"
)

func main() {
	configs.InitDB()

	r := routers.InitRouter()

	addr := fmt.Sprintf("%s:%s", configs.ServerIP, configs.ServerPort)
	r.Run(addr)
}
