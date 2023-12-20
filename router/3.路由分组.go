package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	message string `json:"msg"`
}

func UserList(c *gin.Context) {
	var userInfo []UserInfo = []UserInfo{
		{"fly", 21},
		{"zhangsan", 22},
		{"lisi", 23},
	}
	c.JSON(200, Response{200, userInfo, "success"})
}

func UserRouter(router *gin.RouterGroup) {
	router.GET("/users", UserList)
	router.POST("/users1", UserList)
}

func main() {
	router := gin.Default()

	// 分组, 类似于 springmvn 中每个 controller 的 @RequestMapping
	api := router.Group("/api")
	//{
	//	api.GET("/users", UserList)
	//	api.POST("/users1", UserList)
	//}
	// 抽取到函数中进行路由封装
	UserRouter(api)

	router.Run(":80")
}
