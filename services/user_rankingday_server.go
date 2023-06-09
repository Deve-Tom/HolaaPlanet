package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Ranking_day(ctx *gin.Context) {
	User, _ := strconv.Atoi(ctx.Query("user_id"))
	UserMsg := entity.User{}

	configs.DB.Model(&entity.User{}).Where("user_id = ?", User).First(&UserMsg)
	if UserMsg.UserID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "不存在该用户",
			"ranking":     nil,
		})
		return
	}
	var count int64 = 0
	configs.DB.Model(&entity.User{}).Distinct("day_focus_time").Where("day_focus_time >= ?", UserMsg.DayFocusTime).Count(&count)
	if count < 100 && count > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "查找成功",
			"ranking":     count,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "查找成功",
			"ranking":     "99+",
		})
	}
}
