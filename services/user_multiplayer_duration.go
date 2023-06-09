package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MulDuration(ctx *gin.Context) {
	var req entity.RequestMultiTime
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_msg": "Parsing request body error",
		})
		return
	}

	if ok, time := MulADuration(req.User1ID, req.User2ID, req.Time); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "",
			"loaded":      "success",
			"time":        time,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "token is invalid",
			"loaded":      "fail",
			"time":        0,
		})
	}
}
func MulADuration(User1ID int, User2ID int, Time int) (bool, int) {
	user1 := entity.User{}
	configs.DB.Table("users").Where("user_id = ?", User1ID).First(&user1)

	user2 := entity.User{}
	configs.DB.Table("users").Where("user_id = ?", User2ID).First(&user2)

	//user1
	//日专注时长更新
	//将int类型的Time转换为int64类型
	dayTime := int64(Time) + user1.DayFocusTime
	configs.DB.Table("users").Where("user_id = ?", User1ID).Update("day_focus_time", dayTime)

	//周专注时长更新
	weekTime := int64(Time) + user1.WeekFocusTime
	configs.DB.Table("users").Where("user_id = ?", User1ID).Update("week_focus_time", weekTime)

	//月专注时长更新
	monthTime := int64(Time) + user1.MonthFocusTime
	configs.DB.Table("users").Where("user_id = ?", User1ID).Update("month_focus_time", monthTime)

	//user2专注时长更新
	dayTime2 := int64(Time) + user2.DayFocusTime
	configs.DB.Table("users").Where("user_id = ?", User2ID).Update("day_focus_time", dayTime2)

	weekTime2 := int64(Time) + user2.WeekFocusTime
	configs.DB.Table("users").Where("user_id = ?", User2ID).Update("week_focus_time", weekTime2)

	monthTime2 := int64(Time) + user2.MonthFocusTime
	configs.DB.Table("users").Where("user_id = ?", User2ID).Update("month_focus_time", monthTime2)

	return true, Time
}
