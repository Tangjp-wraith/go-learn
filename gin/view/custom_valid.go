package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 获取结构体中的msg参数
func _GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每个错误信息
			// 根据报错字段获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

func signValid(fl validator.FieldLevel) bool {
	var nameList = []string{"fmy", "gjh", "zhc"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
	router.POST("/", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			// 将err接口断言为具体类型
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, user)
	})

	router.Run("127.0.0.1:8088")
}
