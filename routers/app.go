package routers

import (
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
	apiRouter.POST("/user/register/", services.Register)

	return r
}
