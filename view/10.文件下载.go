package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment;filename="+"技术架构.png")
		c.File("../uploads/技术架构.png")
	})

	router.Run(":80")
}
