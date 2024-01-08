package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	Name string `json:"name" binding:"required"` // 必须有name字段且不为空
	//Name string `json:"name" binding:"contains=f"` // 包含f
	//Name string `json:"name" binding:"excludes=f"` // 不包含f
	//Name string `json:"name" binding:"endswith=f"` // f结尾

	Age int `json:"age" binding:"gt=18,lt=23"`
	//Password   string `json:"password" binding:"min=4,max=6"`
	//RePassword string `json:"re_password" binding:"min=4,max=6"`
	Sex      string   `json:"sex" binding:"oneof=male female"` // 枚举
	LikeList []string `json:"like_list" binding:"required,dive,required,startswith=like"`
	Ip       string   `json:"ip" binding:"ip"`
	Url      string   `json:"url" binding:"url"`
	Uri      string   `json:"uri" binding:"uri"`
	Date     string   `json:"date" binding:"datetime=2006-01-02 15:04:05"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, user)
	})
	router.Run("127.0.0.1:8088")
}
