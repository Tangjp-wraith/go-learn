package main

import (
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int `json:"code"`
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

func _UserListView(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var userList = []UserInfo{
		{"fmy", 21},
		{"zhc", 23},
		{"gjh", 22},
	}
	c.JSON(200, Res{0, userList, "success"})
}

//func MiddleWare(c *gin.Context)  {
//	token := c.GetHeader("token")
//	if token == "1234" {
//		c.Next()
//		return
//	}
//	c.JSON(200, Res{1001, nil, "failed"})
//	c.Abort()
//}

func MiddleWare(msg string) gin.HandlerFunc {
	// 闭包，可以传参
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(200, Res{1001, nil, msg})
		c.Abort()
	}
}

func _UserRouterInit(router *gin.RouterGroup) {
	//userManager := router.Group("user_manager").Use(MiddleWare)
	userManager := router.Group("user_manager").Use(MiddleWare("failed"))
	{
		userManager.GET("/users", _UserListView)
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("api")
	api.GET("/login", func(c *gin.Context) {
		//panic("error hand")
		c.JSON(200, gin.H{"data": "1234"})
	})
	_UserRouterInit(api)

	router.Run("127.0.0.1:8088")
}
