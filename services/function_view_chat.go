package services

import (
	"HolaaPlanet/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// viewFriendRequest
// Maintainers:庹建川 Times:2023-06-10
// Part 1:消息结构体，发送ID，接收ID
// Part 2:获取用户信息服务，如果用户不存在，返回错误，如果成功，返回用户信息
type viewChatRequest struct {
	SendID int `json:"send_user_id" binding:"required"`
}

type ChatListView struct {
	ReceiveID   int    `gorm:"column:receive_user_id" json:"receive_user_id"`
	MessageID   int    `gorm:"column:message_id" json:"message_id"`
	SendMessage string `gorm:"column:send_message" json:"send_message"`
	SendTime    string `gorm:"send_time" json:"send_time"`
	ReadStatus  int    `gorm:"column:read_status" json:"read_status"`
}

// ViewChat
// Maintainers:庹建川 Times:2023-06-13
// Part1: 聊天记录服务处理函数
// Part2: 在viewFriendsByUserID中查询后返回对应的好友列表
func ViewChat(ctx *gin.Context) {
	// 将请求正文绑定到自定义的结构体
	var req viewChatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	chats, err := ViewAChat(req.SendID)
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
		"chats":       chats,
	})
}

// ViewAChat
// Maintainers:庹建川 Times:2023-06-12
// Part1: 在数据库中查询好友的相关信息
func ViewAChat(sendID int) ([]ChatListView, error) {
	// 好友列表数组
	var ChatList []ChatListView
	configs.DB.Table("send_user_messages").Select("send_user_id", "receive_user_id", "message_id", "send_message", "send_time", "read_status").Where("send_user_id = ? ", sendID).Find(&ChatList)
	return ChatList, nil
}
