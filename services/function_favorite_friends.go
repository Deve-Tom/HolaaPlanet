package services

import (
	"HolaaPlanet/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteFriendRequest
// Maintainers:庹建川 Times:2023-06-10
// Part 1:定义了收藏好友的请求格式，用户ID,好友ID，
// Part 2:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
type FavoriteFriendRequest struct {
	UserID   int `json:"user_id" binding:"required"`
	FriendID int `json:"friend_id" binding:"required"`
}

func FavoriteFriend(ctx *gin.Context) {
	// 将请求正文绑定到自定义的结构体
	var req FavoriteFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 返回错误
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return

	}

	if ok, id := FavoriteAFriend(req.UserID, req.FriendID); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "favorite success",
			"friend_id":   id,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "favorite fail",
			"friend_id":   0,
		})
	}
}

// FavoriteAFriend
// 添加字段用update不是Create
func FavoriteAFriend(UserID int, FriendID int) (bool, int) {
	configs.DB.Table("friends_lists").Where("user_id = ? and friend_id = ?", UserID, FriendID).Update("friend_remark", 1)
	return true, FriendID

}
