package routers

import (
	"HolaaPlanet/middleware"
	"HolaaPlanet/services"
	"github.com/gin-gonic/gin"
)

// InitRouter
// Maintainers:贺胜 Times:2023-04-13
// Part 1:初始化路由
// Part 2:初始化路由，包括根路由和API路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 根路由
	apiRouter := r.Group("/holaa")

	// API service
	// Maintainers:贺胜 Times:2023-04-13
	apiRouter.POST("/user/register/", services.Register)
	apiRouter.POST("/user/login/", services.Login)
	apiRouter.GET("/user/", middleware.AuthMiddleWare(), services.UserSingleInfoServer)

	//Maintainers:邵洁  Times:2023-06-08
	apiRouter.PUT("/user/personal_duration/", services.PerDuration)
	apiRouter.PUT("/user/update_information/", services.UpdateInformation)

	return r
}
