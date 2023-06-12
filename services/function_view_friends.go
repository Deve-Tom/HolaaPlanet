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

// FriendListView
// Maintainers:贺胜 Times:2023-06-12
// Part1: 好友信息结构体
// Part2: 用于好友信息查询结果
type FriendListView struct {
	FriendID        int    `gorm:"column:user_id" json:"friend_id"`
	UserAvatar      string `gorm:"column:user_avatar" json:"user_avatar"`
	Nickname        string `gorm:"column:nickname" json:"nickname"`
	PersonSignature string `gorm:"column:person_signature" json:"person_signature"`
}

// ViewFriend
// Maintainers:贺胜 Times:2023-06-12
// Part1: 好友列表服务处理函数
// Part2: 在viewFriendsByUserID中查询后返回对应的好友列表
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
			"status_msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "view success",
		"friends":     friends,
	})
}

// viewFriendsByUserID
// Maintainers:贺胜 Times:2023-06-12
// Part1: 在数据库中查询好友的相关信息
// Part2: 使用子查询的方法将好友的用户数据信息全部查询出来
func viewFriendsByUserID(userID int) ([]FriendListView, error) {
	// 好友列表数组
	var FriendList []FriendListView

	// 错误返回
	subQuery := configs.DB.Table("friends_lists").Select("friend_id").Where("user_id = ?", userID)
	if err := configs.DB.Table("users").Select("user_id", "user_avatar", "nickname", "person_signature").Where("user_id in (?)", subQuery).Find(&FriendList); err != nil {
		return FriendList, err.Error
	}

	// 正常返回
	return FriendList, nil
}
