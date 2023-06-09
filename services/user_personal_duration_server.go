package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PerDuration(ctx *gin.Context) {
	var req entity.RequestPerTime
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_msg": "Parsing request body error",
		})
		return
	}

	if ok, id := PerADuration(req.UserID, req.Time); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "",
			"loaded":      "success",
			"user_id":     id,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg":  "token is invalid",
			"loaded":      "fail",
			"user_id":     0,
		})
	}
}
func PerADuration(UserID int, Time int) (bool, int) {
	user := entity.User{} //创建一个User类型的user对象

	//数据库表中的第一行，赋给user
	configs.DB.Table("users").Where("user_id = ?", UserID).First(&user)

	//日专注时长更新
	//将int类型的Time转换为int64类型
	dayTime := int64(Time) + user.DayFocusTime
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("day_focus_time", dayTime)

	//周专注时长更新
	weekTime := int64(Time) + user.WeekFocusTime
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("week_focus_time", weekTime)

	//月专注时长更新
	monthTime := int64(Time) + user.MonthFocusTime
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("month_focus_time", monthTime)

	return true, UserID
}
