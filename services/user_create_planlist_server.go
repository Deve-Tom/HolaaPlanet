package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreatePlanLists(ctx *gin.Context) {
	var requestBody entity.RequestBodyPlan
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusUnauthorized,
			"status_msg":  "Parsing request body error",
		})
		return
	}

	UserID, _ := strconv.Atoi(requestBody.UserID)
	PlanBeginTime := requestBody.BeginTime
	PlanContent := requestBody.Content

	UserMsg := entity.User{}
	configs.DB.Model(&entity.User{}).Where("user_id = ?", UserID).First(&UserMsg)
	if UserMsg.UserID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":  -1,
			"status_msg":   "不存在该用户",
			"plan_list_id": nil,
		})
		return
	}
	ok, planlistid := CreateAPlan(UserID, PlanBeginTime, PlanContent)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":  0,
			"status_msg":   "success",
			"plan_list_id": planlistid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code":  -1,
			"status_msg":   "fail",
			"plan_list_id": planlistid,
		})
	}
}

func CreateAPlan(UserID int, BeginTime time.Time, Content string) (bool, int) {
	plan := entity.PlanList{
		UserID:        UserID,
		PlanBeginTime: BeginTime,
		PlanContent:   Content,
		PlanListID:    1,
	}
	result := configs.DB.Create(&plan)
	if result != nil {
		log.Print(result.Error)
		return false, 0
	}
	return true, plan.PlanListID
}
