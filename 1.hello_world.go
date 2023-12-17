package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.String(200, "hello world，阿飞")
}

func main() {
	fmt.Println("hello world")
	// 创建一个默认的路由
	router := gin.Default()

	// 绑定路由规则和路由函数, 访问 /index 路由, 将有对应的函数去处理
	router.GET("/index", Index)

	// 启动监听, gin 会把 web 服务运行在本机的 0.0.0.0 的 8080 端口
	router.Run(":8080")

	// 用原生 http 服务的方式, router.Run 本质就是 http.ListenAndServe 的进一步封装
	http.ListenAndServe(":8080", router)
}
