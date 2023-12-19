package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 请求头的各种获取方式
	router.GET("/", func(context *gin.Context) {
		// 字母大小写不区分, 单词与单词之间用 - 连接
		// 用于获取一个请求头
		fmt.Println(context.GetHeader("User-Agent"))
		fmt.Println(context.GetHeader("user-Agent"))
		fmt.Println(context.GetHeader("user-agent"))

		// Header 是一个普通的 Map[string][string]
		fmt.Println(context.Request.Header)
		// 如果是使用 Get 方法或者是 .GetHeader, 那么可以不用区分大小写, 并且返回第一个 value
		fmt.Println(context.Request.Header.Get("User-Agent"))
		// 如果是用 map 的取值方式; 请注意大小写问题
		fmt.Println(context.Request.Header["user-agent"])

		// 自定义请求头
		fmt.Println(context.Request.Header.Get("Token"))
		fmt.Println(context.Request.Header.Get("token"))

		context.JSON(200, gin.H{"msg": "成功"})
	})

	// 爬虫和用户区别对待
	router.GET("/index", func(context *gin.Context) {

	})

	// 设置响应头
	router.GET("/res", func(context *gin.Context) {
		context.Header("token", "seks873dfe")
		context.JSON(200, gin.H{"data": "设置响应头"})
	})

	router.Run(":80")
}
