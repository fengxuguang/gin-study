package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "返回txt")
}

func _json(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 不返回该字段，使用 -
	}
	//user := UserInfo{"fly", 23, "123456"}

	//userMap := map[string]any{
	//	"user_name": "fly",
	//	"age":       23,
	//}

	//c.JSON(200, userMap)
	c.JSON(200, gin.H{"username": "fly", "age": 24})
}

func _xml(c *gin.Context) {
	c.XML(200, gin.H{"username": "fly", "age": 23, "status": http.StatusOK})
}

func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"username": "fly", "age": 23, "status": http.StatusOK})
}

// 响应 html
func _html(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 不返回该字段，使用 -
	}
	user := UserInfo{"fly", 23, "123456"}
	//c.HTML(200, "index.html", gin.H{"username": "fly"})
	c.HTML(200, "index.html", user)
}

// 重定向
func _redirect(c *gin.Context) {
	c.Redirect(302, "http://www.baidu.com")
}

func main() {
	router := gin.Default()

	// 加载模板目录下所有的模板文件
	router.LoadHTMLGlob("templates/*")
	// 在 golang 中, 没有相对文件的路径, 它只有相对项目的路径
	router.StaticFile("/static/hero.jpg", "static/hero.jpg")

	router.GET("/txt", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)

	router.Run(":80")
}
