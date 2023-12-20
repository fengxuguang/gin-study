package main

import "github.com/gin-gonic/gin"

type Res struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}

func _UserList(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var userInfo []UserInfo = []UserInfo{
		{"fly", 21},
		{"zhangsan", 22},
		{"lisi", 23},
	}
	c.JSON(200, Res{200, userInfo, "success"})
}

func _UserRouter(router *gin.RouterGroup) {
	users := router.Group("/users").Use(Middleware)
	users.GET("/users", _UserList)
	users.POST("/users1", _UserList)
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
	_UserRouter(api)
	api.GET("/login", _Middleware("权限校验不成功"), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "123"})
	})

	router.Run(":80")
}

// 闭包, 另一种传参方式
func _Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "123" {
			c.Next()
			return
		}
		c.JSON(200, Res{200, nil, "权限验证失败"})
		c.Abort()
	}
}

func Middleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "123" {
		c.Next()
		return
	}
	c.JSON(200, Res{200, nil, "权限验证失败"})
	c.Abort()
}
