package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SignUserInfo struct {
	Name       string `json:"name" binding:"required,min=4"`          // 用户名
	Age        int    `json:"age"`                                    // 年龄
	Sex        string `json:"sex" binding:"oneof=man women"`          // 性别
	Password   string `json:"password"`                               // 密码
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
	Date       string `json:"date" binding:datetime=2023-12-19`       // 日期
	Url        string `json:"url" binding:"url"`                      // url 是 uri 的子集
	Uri        string `json:"uri" binding:"uri"`                      // uri 是 url 的父集
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
