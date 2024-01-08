package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")
	c.Set("name", "fmy")
	c.Set("user", User{
		Name: "fmy",
		Age:  21,
	})
}

func main() {
	router := gin.Default()
	router.Use(m10)
	router.GET("/m10", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)
		_user, _ := c.Get("user")
		fmt.Println(_user)
		user := _user.(User)
		fmt.Println(user.Name, user.Age)
		c.JSON(200, gin.H{"msg": user})
	})
	router.GET("/m11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "m11"})
	})
	router.Run("127.0.0.1:8088")
}
