package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(200, "hello")
}

// 响应json 重点
func _json(c *gin.Context) {
	// json响应结构体
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 忽略转换为json
	}
	user := UserInfo{"fmy", 22, "123456"}
	c.JSON(200, user)

	//json响应map
	//userMap := map[string]string{
	//	"user_name": "fmy",
	//	"age":       "23",
	//}
	//c.JSON(200, userMap)

	// 直接响应json
	//c.JSON(200, gin.H{"username": "fmy", "age": 22})
}

// 响应xml
func _xml(c *gin.Context) {
	c.XML(200, gin.H{"user": "fmy", "message": "sb", "status": http.StatusOK, "data": gin.H{"user": "fmy1"}})
}

// 响应yaml
func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"user": "fmy", "message": "sb", "status": http.StatusOK, "data": gin.H{"user": "fmy1"}})
}

// 响应html
func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "fmy"})
}

// 重定向
func _redirect(c *gin.Context) {
	//c.Redirect(301, "https://www.baidu.com")
	//c.Redirect(302, "https://www.fengfengzhidao.com")
	c.Redirect(301, "/html")
}

func main() {
	router := gin.Default()
	// 加载模板目录下所有的模板文件
	router.LoadHTMLGlob("templates/*")
	// 在项目中，没有相对于文件的路径，只有相对于项目的路径
	// 配置单个文件，网页请求的路由，文件的路径
	router.StaticFile("/pzw", "static/pzw.jpg")
	// 网页请求这个静态目录的前缀，第二个参数是一个目录，前缀不重要
	router.StaticFS("/static", http.Dir("static/static"))

	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)

	router.Run("127.0.0.1:8088")
}
