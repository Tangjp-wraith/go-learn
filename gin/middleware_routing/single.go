package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 ... in")
	//c.JSON(200, gin.H{"msg": "m1"})
	c.Abort()
	c.Next()
	fmt.Println("m1 ... out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 ... in")
	//c.JSON(200, gin.H{"msg": "m2"})
	c.Next()
	fmt.Println("m2 ... out")
}

func main() {
	router := gin.Default()

	router.GET("/", m1, func(c *gin.Context) {
		fmt.Println("index ... in")
		c.JSON(200, gin.H{"msg": "index"})
		c.Next()
		fmt.Println("index ... out")
	}, m2)

	router.Run("127.0.0.1:8088")
}
