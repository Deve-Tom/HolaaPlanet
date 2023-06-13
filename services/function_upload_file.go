package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// HandleUploadFile
// Maintainers:贺胜 Times:2023-06-11
// Part 1:用于更新用户头像
// Part 2:更新用户头像，并将其写入。/static/avatar/xxx.png之中，同时更新数据库
func HandleUploadFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	filename := ctx.PostForm("filename")

	if filename == "" {
		filename = file.Filename
	}
	err := ctx.SaveUploadedFile(file, "./static/user_avatar/"+filename)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err:%s", err))
		return
	}

	// 更新数据库的头像信息
	var User entity.User
	User.UserID, _ = strconv.Atoi(filename[:len(filename)-4])
	UserMsg := "./static/user_avatar/" + filename
	configs.DB.Model(&entity.User{}).Where("user_id = ?", User).Update("user_avatar", UserMsg)

	ctx.String(http.StatusOK, fmt.Sprintf("upload success"))
}
