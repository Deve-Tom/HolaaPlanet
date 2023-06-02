package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AddFriend(ctx *gin.Context) {
	UserID, _ := strconv.Atoi(ctx.Query("user_id"))
	FriendID, _ := strconv.Atoi(ctx.Query("friend_id"))

	//configs.DB.Model(&entity.User{}).Where("user_id = ?", User).First(&UserMsg)
	if ok, id := CreateAFriend(UserID, FriendID); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "Add_success",
			"friend_id":   id,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "Add_fail",
			"friend_id":   -1,
		})
	}
}
func CreateAFriend(UserID int, FriendID int) (bool, int) {

	result := configs.DB.Find(&entity.AddFriends{})

	if result.Error != nil {
		log.Print(result.Error)
		return false, 0
	} else {
		if result.RowsAffected == 0 {
			user := entity.AddFriends{
				UserID:   UserID,
				FriendID: FriendID,
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false, 0
			}
		} else {
			user := entity.AddFriends{
				UserID:   UserID,
				FriendID: FriendID,
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false, 0
			}
		}
	}
	return true, FriendID
}
