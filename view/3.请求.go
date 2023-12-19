package main

/**
四大请求方式
Restful 风格指的是网络应用中就是资源定位和资源操作的风格, 不是标准也不是协议
GET: 从服务器取出资源(一项或多项)
POST: 在服务器新建一个资源
PUT: 在服务器更新资源(客户端提供完整资源数据)
PATCH: 在服务器更新资源(客户端提供需要修改的资源数据)
DELETE: 从服务器删除资源
*/

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.GET("/articles", _getArticles)
	router.GET("/articles/:id", _getArticleDetails)

	router.Run(":80")
}

type ArticleModel struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

var articles []ArticleModel = []ArticleModel{
	{1, "Java入门", "这是Java入门相关"},
	{2, "Go入门", "这是Go入门相关"},
	{3, "Gin入门", "这是Gin入门相关"},
}

// 获取文章列表
func _getArticles(c *gin.Context) {
	c.JSON(200, Response{200, articles, "success"})
}

// 获取文章数据
func _getArticleDetails(c *gin.Context) {
	articleId := c.Param("id")

	for i := 0; i < len(articles); i++ {
		atoi, _ := strconv.Atoi(articleId)
		if atoi == articles[i].Id {
			c.JSON(200, Response{200, articles[i], "查询成功"})
			return
		}
	}
	c.JSON(400, Response{400, "", "查询失败"})
}

// 直接查询
func _query(c *gin.Context) {
	username := c.Query("username")
	query, b := c.GetQuery("username")
	fmt.Println("username=", username)
	fmt.Println("query=", query)
	fmt.Println("b=", b)
}

// json 转结构体方法
func bindJson(context *gin.Context, obj any) (err error) {
	body, _ := context.GetRawData()
	contentType := context.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

// 解析原始参数
func _raw(context *gin.Context) {
	type User struct {
		Name string
		Age  int
	}
	var user User
	bindJson(context, &user)
	fmt.Println(user)
	//body, _ := context.GetRawData()
	//contentType := context.GetHeader("Content-Type")
	//switch contentType {
	//case "application/json":
	//	type User struct {
	//		Name string
	//		Age  int
	//	}
	//	var user User
	//	err := json.Unmarshal(body, &user)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(user)
	//}
}

// 表单 postform
func _form(context *gin.Context) {
	fmt.Println(context.PostForm("name"))
	fmt.Println(context.PostFormArray("name"))
}

// 动态参数
func _param(context *gin.Context) {
	fmt.Println(context.Param("user_id"))
	fmt.Println(context.Param("book_id"))
}
