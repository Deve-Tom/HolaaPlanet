package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AddFriendRequest
// Maintainers:庹建川 Times:2023-06-10
// Part 1:定义了好友添加的请求格式，用户ID,好友ID，
// Part 2:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
type AddFriendRequest struct {
	UserID   int `json:"user_id" binding:"required"`
	FriendID int `json:"friend_id" binding:"required"`
}

func AddFriend(ctx *gin.Context) {
	// 将请求正文绑定到自定义的结构体
	var req AddFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 返回错误
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return

	}

	// 调用 CreateAFriend 函数，处理请求
	if ok, id := CreateAFriend(req.UserID, req.FriendID); ok {
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

// 用PUT请求方法返回

func CreateAFriend(UserID int, FriendID int) (bool, int) {

	result := configs.DB.Find(&entity.AddFriends{})
	result1 := configs.DB.Find(&entity.FriendsList{})
	//如果已经存在将多余的删除掉在friends_list表中（因为friends_lists表中添加好友会重复出现，所以删除多余的）
	user := entity.FriendsList{}
	configs.DB.Table("friends_lists").Where("user_id = ? and friend_id = ?", UserID, FriendID).Delete(&user)

	if result1.Error != nil {
		log.Print(result.Error)
		return false, 0
	} else {
		if result1.RowsAffected == 0 {
			user := entity.FriendsList{
				UserID:   UserID,
				FriendID: FriendID,
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false, 0
			}
		} else {
			user := entity.FriendsList{
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
