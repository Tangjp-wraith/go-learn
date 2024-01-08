package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
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

// 文章列表页面
func _getList(c *gin.Context) {
	articleList := []ArticleModel{
		{"golang learning", "golang is xxx"},
		{"java learning", "java is xxx"},
		{"cpp learning", "cpp is xxx"},
	}
	c.JSON(200, Response{0, articleList, "success"})
}

// 文章详情
func _getDetail(c *gin.Context) {
	// 获取param中的id
	fmt.Println(c.Param("id"))
	article := []ArticleModel{
		{"golang learning", "golang is xxx"},
	}
	c.JSON(200, Response{0, article, "success"})
}

// 创建文章
func _create(c *gin.Context) {
	// 接收前端传来的json数据
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "add success"})
}

// 编辑文章
func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "edit success"})
}
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{0, map[string]string{}, "delete success"})
}

func main() {
	router := gin.Default()
	router.GET("/articles", _getList)       //文章列表
	router.GET("/articles/:id", _getDetail) //文章详情
	router.POST("/articles", _create)       //添加文章
	router.PUT("/articles/:id", _update)    //编辑文章
	router.DELETE("/articles/:id", _delete) //删除文章

	router.Run("127.0.0.1:8088")
}
