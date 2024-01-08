package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认路由
	router := gin.Default()
	// 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "fuck fmy")
	})
	// 启动监听，gin会把web服务运行在本机0.0.0.0.8088
	router.Run(":8088")
	// 用原生http服务的方式，router.Run本质就是http.ListenAndServe的进一步封装
	//http.ListenAndServe(":8088", router)
}
