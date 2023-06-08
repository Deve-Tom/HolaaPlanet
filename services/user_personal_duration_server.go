package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type PerTime struct {
	UserID    int    `gorm:"column:user_id"`
	DayTime   string `gorm:"column:day_focus_time"`
	WeekTime  string `gorm:"column:week_focus_time"`
	MonthTime string `gorm:"column:month_focus_time"`
}

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
func PerADuration(UserID int, Time string) (bool, int) {
	user := entity.User{} //创建一个User类型的user对象

	//数据库表中的第一行，赋给user
	configs.DB.Table("users").Where("user_id = ?", UserID).First(&user)

	//将string类型的Time转换为int类型
	onceTime, err := strconv.Atoi(Time)
	if err != nil {
		fmt.Println("Invalid Time...")
		return false, 0
	}

	//日专注时长更新
	//将time.Time类型的user.DayFocusTime转换为int64类型(时间转换为以秒为单位的值)
	s1 := user.DayFocusTime.UnixNano() / int64(time.Second)

	//将s1强制类型转换int64->int，并且做加法赋值给time1
	time1 := onceTime + int(s1)

	// 将时间戳转换为 time.Time 对象
	timestamp1 := time.Unix(int64(time1), 0)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("day_focus_time", timestamp1)

	//周专注时长更新
	s2 := user.WeekFocusTime.UnixNano() / int64(time.Second)
	time2 := onceTime + int(s2)
	timestamp2 := time.Unix(int64(time2), 0)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("week_focus_time", timestamp2)

	//月专注时长更新
	s3 := user.MonthFocusTime.UnixNano() / int64(time.Second)
	time3 := onceTime + int(s3)
	timestamp3 := time.Unix(int64(time3), 0)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("month_focus_time", timestamp3)

	return true, UserID
}
