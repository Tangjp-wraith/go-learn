package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList = []UserInfo{
		{"fmy", 21},
		{"zhc", 23},
		{"gjh", 22},
	}
	c.JSON(200, Response{0, userList, "success"})
}

func ArticleListView(c *gin.Context) {
	var ArticleList = []ArticleInfo{
		{"go-book", "go is xxx"},
		{"rust-book", "rust is xxx"},
	}
	c.JSON(200, Response{0, ArticleList, "success"})
}

func UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	{
		userManager.GET("/users", UserListView)
	}
}

func ArticleRouterInit(router *gin.RouterGroup) {
	articleManager := router.Group("article_manager")
	{
		// api/article_manager/articles
		articleManager.GET("/articles", ArticleListView)
	}
}

func main() {
	router := gin.Default()

	api := router.Group("api")

	UserRouterInit(api)
	ArticleRouterInit(api)

	router.Run("127.0.0.1:8088")
}
