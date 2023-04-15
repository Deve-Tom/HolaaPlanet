package services

import (
	"HolaaPlanet/common"
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Login
// Maintainers:贺胜 Times:2021-04-15
// Part 1:用户登陆
// Part 2:用户登陆服务，如果用户不存在，返回错误，如果密码错误，返回错误，如果成功，返回token
func Login(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	password := ctx.Query("password")

	// 用户认证
	token, err := UserAuthServer(userId, password)

	// 用户认证错误处理
	if err != nil {
		if entity.ERROR_USER.ErrorToString(err) == entity.ERROR_USER.UserNotFound.Error() {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "用户不存在",
			})
			return
		} else if entity.ERROR_USER.ErrorToString(err) == entity.ERROR_USER.UserPasswordError.Error() {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "密码错误",
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "未知错误",
			})
			return
		}
	}

	// 用户认证成功
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  "登录成功",
		"token":       token,
	})
}

// UserAuthServer
// Maintainers:贺胜 Times:2021-04-15
// Part 1:用户认证
// Part 2:用户认证服务，如果用户不存在，返回错误，如果密码错误，返回错误，如果成功，返回token
func UserAuthServer(userId int, password string) (string, error) {
	user := entity.User{}
	configs.DB.Table("users").Where("user_id = ?", userId).First(&user)

	if user.UserID == 0 {
		return "", entity.ERROR_USER.UserNotFound
	} else if user.Password != password {
		return "", entity.ERROR_USER.UserPasswordError
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		return "", err
	}

	// 更新token
	configs.DB.Model(&user).Where("user_id = ?", user.UserID).Update("user_token", token)

	return token, nil
}
