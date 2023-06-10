package services

import (
	"HolaaPlanet/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// viewFriendRequest
// Maintainers:庹建川 Times:2023-06-10
// Part 1:好友结构体，好友ID
// Part 2:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
type viewFriendRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

type friend struct {
	FriendID int `json:"friend_id"`
}

type friendTemp struct {
	FriendID string `gorm:"column:group_concat(friend_id)"`
}

func ViewFriend(ctx *gin.Context) {
	// 将请求正文绑定到自定义的结构体
	var req viewFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	friends, err := viewFriendsByUserID(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "view fail",
			"friends":     []friend{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "view success",
		"friends":     friends,
	})
}

// 利用user_id分组，传出friends（string）
func viewFriendsByUserID(userID int) ([]friendTemp, error) {
	var friends []friendTemp
	configs.DB.Table("friends_lists").Select("group_concat(friend_id)").Group("user_id").Having("user_id = ?", userID).Scan(&friends)
	return friends, nil
}
