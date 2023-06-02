package services

import (
	"HolaaPlanet/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteFriend(ctx *gin.Context) {
	UserID, _ := strconv.Atoi(ctx.Query("user_id"))
	FriendID, _ := strconv.Atoi(ctx.Query("check_friend_id"))

	if ok, id := FavoriteAFriend(UserID, FriendID); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":     0,
			"status_msg":      "favorite success",
			"check_friend_id": id,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":     -1,
			"status_msg":      "favorite fail",
			"check_friend_id": 0,
		})
	}
}

func FavoriteAFriend(UserID int, FriendID int) (bool, int) {
	configs.DB.Table("friends_lists").Where("user_id = ? and friend_id = ?", UserID, FriendID).Select("friend_remark").Create(1)
	return true, FriendID
}
