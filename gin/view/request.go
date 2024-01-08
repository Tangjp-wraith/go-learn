package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询参数
func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	// 拿到多个相同的查询参数
	fmt.Println(c.QueryArray("user"))
	fmt.Println(c.DefaultQuery("addr", "JiangSu"))
	// ?user[name]="fmy"&user[age]=12 return map[age:12 name:"fmy"]
	//fmt.Println(c.QueryMap("user"))
}

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "JiangSu")) // 如果用户没传就使用默认值
	forms, err := c.MultipartForm()                   // 接受所有form参数，包括文件
	fmt.Println(forms, err)
}

func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			return err
		}
	}
	return nil
}

// 原始参数
func _raw(c *gin.Context) {
	//body, _ := c.GetRawData()
	//contentType := c.GetHeader("Content-Type")
	//switch contentType {
	//case "application/json":
	//	type User struct {
	//		Name string `json:"name"`
	//		Age  int    `json:"age"`
	//	}
	//	var user User
	//	err := json.Unmarshal(body, &user)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(user)
	//}
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)

	router.Run("127.0.0.1:8088")

}
