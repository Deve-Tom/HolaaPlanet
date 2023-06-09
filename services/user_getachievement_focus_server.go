package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAchievement_focus(ctx *gin.Context) {
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
	var times int64
	times = UserMsg.MonthFocusTime
	var minutes = times / 60
	var title int64
	switch minutes {
	case 10:
		title = 1
		break
	case 100:
		title = 2
		break
	case 200:
		title = 3
		break
	case 400:
		title = 4
		break
	case 600:
		title = 5
		break
	case 800:
		title = 6
		break
	case 1800:
		title = 7
		break
	case 3500:
		title = 8
		break
	case 6000:
		title = 9
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
