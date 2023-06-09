package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func UpdateInformation(ctx *gin.Context) {
	var req entity.RequestUpdateInformation
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_msg": "Parsing request body error",
		})
		return
	}

	if ok, id := UpdateAInformation(req.UserID, req.NickName, req.Signature, req.PassWord); ok {
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
func UpdateAInformation(UserID int, NickName string, Signature string, PassWord string) (bool, int) {
	user := entity.User{} //创建一个User类型的user对象

	//数据库表中的第一行，赋给user
	configs.DB.Table("users").Where("user_id = ?", UserID).First(&user)

	user.Nickname = NickName
	user.Password = PassWord
	user.Signature = Signature
	user.UserAvatar = fmt.Sprintf("./static/user_avatar/%d.png", user.UserID)
	_, err := os.Stat(user.UserAvatar)
	if err != nil || os.IsNotExist(err) {
		configs.DB.Table("users").Where("user_id = ?", UserID)
	}
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("nickname", user.Nickname)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("password", user.Password)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("person_signature", user.Signature)
	configs.DB.Table("users").Where("user_id = ?", UserID).Update("user_avatar", user.UserAvatar)
	return false, 0
}
