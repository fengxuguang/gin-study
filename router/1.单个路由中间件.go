package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 ...")
	c.JSON(200, gin.H{"msg": "m1 response"})
	// c.Abort() 拦截, 后续方法不执行
	//c.Abort()
	c.Next()
	fmt.Println("m1 ... out")
}

func m2(c *gin.Context) {
	fmt.Println("m2 ...")
	c.Next()
	fmt.Println("m2 ... out")
}

func main() {
	router := gin.Default()

	router.GET("/", m1, func(c *gin.Context) {
		fmt.Println("index ...")
		c.JSON(200, gin.H{"msg": "成功"})
		c.Next()
		c.Abort()
		fmt.Println("index ... out")
	}, m2)

	router.Run(":80")
}
