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

	// Maintainers:陈微雨 Times:2023-06-09
	apiRouter.GET("/user/send_message/", middleware.AuthMiddleWare(), services.Handler)

	// Maintainers:宋昭城 Times:2023-06-03
	apiRouter.GET("/user/ranking_day/", middleware.AuthMiddleWare(), services.Ranking_day)
	apiRouter.GET("/user/ranking_week/", middleware.AuthMiddleWare(), services.Ranking_week)
	apiRouter.GET("/user/ranking_month/", middleware.AuthMiddleWare(), services.Ranking_month)
	apiRouter.POST("/user/get_achievements_focus/", middleware.AuthMiddleWare(), services.GetAchievement_focus)
	apiRouter.POST("/user/get_achievements_friends/", middleware.AuthMiddleWare(), services.GetAchievement_friends)
	apiRouter.POST("/user/get_achievements_plans/", middleware.AuthMiddleWare(), services.GetAchievement_plans)
	apiRouter.POST("/user/get_Quotes/", middleware.AuthMiddleWare(), services.GetQuotes)
	apiRouter.PUT("/user/create_plan_lists/", middleware.AuthMiddleWare(), services.CreatePlanLists)

	//Maintainers:邵洁  Times:2023-06-08
	apiRouter.PUT("/user/personal_duration/", services.PerDuration)
	apiRouter.POST("/user/update_information/", services.UpdateInformation)
	apiRouter.PUT("/user/multiplayer_duration/", services.MulDuration)

	return r
}
