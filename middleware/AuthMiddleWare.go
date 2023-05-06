package middleware

import (
	"HolaaPlanet/common"
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleWare
// Maintainers:贺胜 Times:2023-04-14
// Part 1:验证token
// Part 2:验证token是否为空
// Part 3:验证token是否有效
// Part 4:验证用户是否存在
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token和user_id
		tokenString := ctx.Query("token")
		userID := ctx.Query("user_id")

		// 验证token是否为空
		if userID == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "user_id is empty",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证token是否为空
		if tokenString == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "token is empty",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证token是否有效
		parseToken, claims, err := common.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "token is invalid",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证用户是否存在
		var User entity.User
		configs.DB.Model(&entity.User{}).Where("user_id = ?", claims.UserID).First(&User)

		// 用户不存在
		if User.UserID == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "user is not exist",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证token无效
		if !parseToken.Valid {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "token is invalid",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证user_id是否有效
		if User.UserID != claims.UserID {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": -1,
				"status_msg":  "user_id is invalid",
				"user":        nil,
			})
			ctx.Abort()
			return
		}

		// 验证成功
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
