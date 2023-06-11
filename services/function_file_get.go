package services

import "github.com/gin-gonic/gin"

// HandleDownloadFile
// Maintainers:贺胜 Times:2023-06-11
// Part 1:用于获取用户所需资源
// Part 2:前往对应目录查找文件并返回
func HandleDownloadFile(c *gin.Context) {
	filepath := c.Query("content")

	c.File(filepath)
}
