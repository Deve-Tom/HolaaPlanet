package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteFriendRequest 定义了删除好友请求的格式，用户，好友ID
// Part 1:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
type DeleteFriendRequest struct {
	UserID   int `json:"user_id" binding:"required"`   // 要删除好友的用户 ID，必填字段
	FriendID int `json:"friend_id" binding:"required"` // 要删除的好友 ID，必填字段
}

// DeleteFriend 处理删除好友的请求
func DeleteFriend(ctx *gin.Context) {
	// 将请求正文绑定到自定义的结构体
	var req DeleteFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 如果请求不完整或格式不正确，返回一个有意义的错误响应
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, id := DeleteAFriend(req.UserID, req.FriendID); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "Delete_success",
			"friend_id":   id,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "Delete_fail",
			"friend_id":   -1,
		})
	}

}

// DeleteAFriend
// Add_friend表和Friends_lists表都要删除
func DeleteAFriend(UserID int, FriendID int) (bool, int) {
	user := entity.AddFriends{}
	configs.DB.Table("add_friends").Where("user_id = ? and friend_id = ?", UserID, FriendID).Delete(&user)
	user2 := entity.FriendsList{}
	configs.DB.Table("friends_lists").Where("user_id = ? and friend_id = ?", UserID, FriendID).Delete(&user2)
	return true, FriendID

}
