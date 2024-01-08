package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()
	// 请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		// 使用.GetHeader或者.Header.Get不区分大小写，并且返回第一个value
		fmt.Println(c.GetHeader("User-Agent"))
		// Header是一个普通的 map[string][]string
		fmt.Println(c.Request.Header.Get("User-Agent"))
		// 区分大小写，可以取到重复的全部value
		fmt.Println(c.Request.Header["User-Agent"])
		// 自定义请求头
		fmt.Println(c.Request.Header.Get("Token"))
		c.JSON(200, gin.H{"msg": "success"})
	})
	// 爬虫和用户区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		// 用正则去匹配
		// 字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬虫
			c.JSON(0, gin.H{"data": "sb"})
			return
		}
		c.JSON(0, gin.H{"data": "fmy"})
	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("token", "tjp run after gyc")
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(200, gin.H{"data": "success"})
	})

	router.Run("127.0.0.1:8088")
}
