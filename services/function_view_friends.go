package services

import (
	"HolaaPlanet/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	//friendList := make([]friend, len(friends))
	//for i, f := range friends {
	//	friendList[i] = friend{FriendID: f.FriendID}
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "view success",
		"friends":     friends,
	})
}

func viewFriendsByUserID(userID int) ([]friendTemp, error) {
	var friends []friendTemp
	configs.DB.Table("friends_lists").Select("group_concat(friend_id)").Group("user_id").Having("user_id = ?", userID).Scan(&friends)
	return friends, nil
}
