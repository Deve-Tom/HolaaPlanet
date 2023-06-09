package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAchievement_friends(ctx *gin.Context) {
	var requestBody entity.RequestBodyGets
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusUnauthorized,
			"status_msg":  "Parsing request body error",
		})
		return
	}
	userId, _ := strconv.Atoi(requestBody.UserID)
	UserMsg := entity.User{}
	configs.DB.Model(&entity.User{}).Where("user_id = ?", userId).First(&UserMsg)
	if UserMsg.UserID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":  -1,
			"status_msg":   "不存在该用户",
			"achievements": nil,
		})
		return
	}
	var count int64
	configs.DB.Model(&entity.FriendsList{}).Where("user_id = ?", userId).Count(&count)
	var title int64
	switch count {
	case 3:
		title = 10
		break
	case 15:
		title = 11
		break
	case 50:
		title = 12
		break
	default:
		title = 0
		break
	}
	achievementMsg := entity.AchievementList{}
	if title == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":  0,
			"status_msg":   "未获得新成就",
			"achievements": nil,
		})
	} else {
		configs.DB.Model(&entity.AchievementList{}).Where("achievement_id = ?", title).First(&achievementMsg)
		getachievement := entity.GetAchievement{
			GetUserID:        UserMsg.UserID,
			GetAchievementID: achievementMsg.AchievementID,
			GetTime:          time.Now(),
			GetStatus:        1,
		}
		insert := configs.DB.Create(&getachievement)
		if insert.Error != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code":  0,
				"status_msg":   "已获得该成就，不可重复获得",
				"achievements": nil,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "达成成就",
				"achievements": struct {
					AchievementTitle   string `json:"achievement_title"`
					AchievementContact string `json:"achievement_contact"`
				}{
					AchievementTitle:   achievementMsg.AchievementTitle,
					AchievementContact: achievementMsg.AchievementContact,
				},
			})
			return
		}
	}
}
