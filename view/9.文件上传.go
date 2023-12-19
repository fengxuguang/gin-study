package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()

	// 单个文件上传
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		fmt.Println(file.Size / 1024) // 单位是字节
		//c.SaveUploadedFile(file, "../uploads/"+file.Filename)

		// Create + Copy
		readerFile, _ := file.Open()
		//data, _ := io.ReadAll(readerFile)
		//fmt.Println(string(data))
		write, _ := os.Create("../uploads/" + file.Filename)
		defer write.Close()
		n, _ := io.Copy(write, readerFile)
		fmt.Println(n)

		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	// 多文件上传
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "../uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传 %d 个文件", len(files))})
	})

	router.Run(":80")
}
