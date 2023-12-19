package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 返回结构体中的 msg 参数
func getValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	// 将 err 接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			if f, exists := getObj.Elem().FieldByName(e.Field()); exists {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"用户名不能为空"`
			Age  int    `json:"age" binding:"required" msg:"年龄参数错误"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": getValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
