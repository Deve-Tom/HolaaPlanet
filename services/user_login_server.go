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
// Maintainers:贺胜 Times:2023-05-20
// Part 1:用户登陆
// Part 2:用户登陆服务，如果用户不存在，返回错误，如果密码错误，返回错误，如果成功，返回token
// BUG: 紧急Bug修复，解决使用Json文件进行数据交流时，无法识别的情况
func Login(ctx *gin.Context) {
	var requestBody entity.RequestBodyUserLogin
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusUnauthorized,
			"status_msg":  "Parsing request body error",
		})
		return
	}

	userId, _ := strconv.Atoi(requestBody.UserID)
	password := requestBody.Password
	// 用户认证
	token, err := UserAuthServer(userId, password)

	// 用户认证错误处理
	if err != nil {
		if entity.ErrorUser.ErrorToString(err) == entity.ErrorUser.UserNotFound.Error() {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "user not exist",
			})
			return
		} else if entity.ErrorUser.ErrorToString(err) == entity.ErrorUser.UserPasswordError.Error() {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "password error",
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusUnauthorized,
				"status_msg":  "unknown error",
			})
			return
		}
	}

	// 用户认证成功
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  "login successfully",
		"token":       token,
	})
}

// UserAuthServer
// Maintainers:贺胜 Times:2023-04-15
// Part 1:用户认证
// Part 2:用户认证服务，如果用户不存在，返回错误，如果密码错误，返回错误，如果成功，返回token
func UserAuthServer(userId int, password string) (string, error) {
	user := entity.User{}
	configs.DB.Table("users").Where("user_id = ?", userId).First(&user)

	if user.UserID == 0 {
		return "", entity.ErrorUser.UserNotFound
	} else if user.Password != password {
		return "", entity.ErrorUser.UserPasswordError
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		return "", err
	}

	// 更新token
	configs.DB.Model(&user).Where("user_id = ?", user.UserID).Update("user_token", token)

	return token, nil
}
