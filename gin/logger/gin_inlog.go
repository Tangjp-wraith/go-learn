package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//file, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf(
	//		"[ fmy ] %v %v %v %v \n",
	//		httpMethod,
	//		absolutePath,
	//		handlerName,
	//		nuHandlers,
	//	)
	//}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"[FMY] %s |%d| %s %s\n",
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			params.StatusCode,
			params.Method,
			params.Path,
		)
	}))

	router.GET("/index", func(c *gin.Context) {})
	router.POST("/users", func(c *gin.Context) {})
	router.PUT("/articles", func(c *gin.Context) {})
	router.DELETE("/articles/:id", func(c *gin.Context) {})

	//fmt.Println(router.Routes())

	//for _, info := range router.Routes() {
	//	fmt.Println(info.Path, info.Method, info.Handler)
	//}

	router.Run("127.0.0.1:8088")
}
