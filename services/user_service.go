package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"HolaaPlanet/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// Register
// Maintainers:贺胜 Times:2023-04-13
// Part 1:用户注册
// Part 2:用户注册，注册成功则返回用户ID和用户Token，否则返回0和空字符串
func Register(ctx *gin.Context) {
	userName := ctx.Query("username")
	password := ctx.Query("password")
	personSignature := ctx.Query("signature")

	if ok, id, userToken := CreateAUser(userName, password, personSignature); ok {
		ctx.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "success",
			"user_id":     id,
			"token":       userToken,
		})
	} else {
		ctx.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "fail",
			"user_id":     0,
			"token":       "",
		})
	}
}

// CreateAUser
// Maintainers:贺胜 Times:2023-04-13
// Part 1:创建用户
// Part 2:创建用户，创建成功则返回true和用户ID和用户Token，否则返回false和0和空字符串
func CreateAUser(userName string, password string, personSignature string) (bool, int, string) {
	result := configs.DB.Find(&entity.User{})

	if result.Error != nil {
		log.Print(result.Error)
		return false, 0, ""
	} else {
		if result.RowsAffected == 0 {
			user := entity.User{
				Nickname:  userName,
				Password:  password,
				Signature: personSignature,
				UserID:    10012,
				UserToken: utils.GetSha256(userName + password),
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false, 0, ""
			}
		} else {
			user := entity.User{
				Nickname:  userName,
				Password:  password,
				Signature: personSignature,
				UserToken: utils.GetSha256(userName + password),
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false, 0, ""
			}
		}
	}

	user := entity.User{}
	configs.DB.Table("users").Where("nickname = ? and password = ?", userName, password).First(&user)

	return true, user.UserID, user.UserToken
}
