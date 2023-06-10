package main

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/routers"
	"HolaaPlanet/services"
	"fmt"
)

func main() {
	configs.InitDB()
	go services.Manager.Start()
	r := routers.InitRouter()

	addr := fmt.Sprintf("%s:%s", configs.ServerIP, configs.ServerPort)
	r.Run(addr)
}
