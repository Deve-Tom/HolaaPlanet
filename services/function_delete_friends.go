package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteFriend(ctx *gin.Context) {
	UserID, _ := strconv.Atoi(ctx.Query("user_id"))
	FriendID, _ := strconv.Atoi(ctx.Query("friend_id"))

	if ok, id := DeleteAFriend(UserID, FriendID); ok {
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
func DeleteAFriend(UserID int, FriendID int) (bool, int) {
	user := entity.AddFriends{}
	configs.DB.Table("add_friends").Where("user_id = ? and friend_id = ?", UserID, FriendID).Delete(&user)
	user2 := entity.FriendsList{}
	configs.DB.Table("friends_lists").Where("user_id = ? and friend_id = ?", UserID, FriendID).Delete(&user2)
	return true, FriendID
}
