package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// UserSingleInfoServer
// Maintainers:贺胜 Times:2021-04-15
// Part 1:获取用户信息
// Part 2:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
func UserSingleInfoServer(ctx *gin.Context) {
	User, _ := strconv.Atoi(ctx.Query("user_id"))
	UserMsg := entity.User{}

	configs.DB.Model(&entity.User{}).Where("user_id = ?", User).First(&UserMsg)
	if UserMsg.UserID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "用户不存在",
			"user":        nil,
		})
		return
	}

	if UserMsg.UserAvatar == "" {
		UserMsg.UserAvatar = fmt.Sprintf("./static/user_avatar/default.png")
		configs.DB.Model(&entity.User{}).Where("user_id = ?", User).Update("user_avatar", UserMsg.UserAvatar)
	}

	theUserAvatar := fmt.Sprintf("./static/user_avatar/%d.png", UserMsg.UserID)
	_, err := os.Stat(theUserAvatar)
	if err != nil || os.IsNotExist(err) {
		UserMsg.UserAvatar = fmt.Sprintf("./static/user_avatar/default.png")
		configs.DB.Model(&entity.User{}).Where("user_id = ?", User).Update("user_avatar", UserMsg.UserAvatar)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"status_msg":  "获取成功",
		"user": struct {
			UserID    int    `json:"user_id"`
			NickName  string `json:"nickname"`
			Avatar    string `json:"avatar"`
			Signature string `json:"signature"`
		}{
			UserID:    UserMsg.UserID,
			NickName:  UserMsg.Nickname,
			Avatar:    UserMsg.UserAvatar,
			Signature: UserMsg.Signature,
		},
	})
}
