package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/download", func(c *gin.Context) {
		// 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Type", "application/octet-stream")
		// 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename="+"pzw.png")
		c.Header("msg", "download success")
		c.File("uploads/pzw.jpg")
	})

	router.Run("127.0.0.1:8088")
}
