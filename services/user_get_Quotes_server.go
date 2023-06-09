package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
	"strconv"
)

func GetQuotes(ctx *gin.Context) {
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
			"status_code": -1,
			"status_msg":  "不存在该用户",
			"Quotes":      nil,
		})
		return
	}
	var ptext []string
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnHTML("div.txt#txt", func(h *colly.HTMLElement) {
		h.ForEach("p", func(i int, h *colly.HTMLElement) {
			// 获取标签内嵌的文本
			ptext = append(ptext, h.Text)
		})
	})
	err = c.Visit("https://www.mingyannet.com/mingyan/228649546")
	randomIndex := rand.Intn(len(ptext))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "获取失败",
			"Quotes":      ptext[randomIndex],
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "获得名言警句",
			"Quotes":      ptext[randomIndex],
		})
	}
}
