package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m10(c *gin.Context) {
	fmt.Println("m10 ...")
	// 中间件之间通过 c.Set(key, value) 进行数据传递
	c.Set("name", "fly")
}

func main() {
	router := gin.Default()

	// 设置全局路由
	router.Use(m10)

	router.GET("/", func(c *gin.Context) {
		name, exists := c.Get("name")
		if exists {
			fmt.Println(name)
		}
		c.JSON(200, gin.H{"msg": "index"})
	})
	router.GET("m11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "m11"})
	})

	router.Run(":80")
}
