package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		//fmt.Println(file.Filename)
		//fmt.Println(file.Size) // 单位是字节
		//c.SaveUploadedFile(file, "./uploads/pzw.png")

		//readerFile, _ := file.Open()
		//data, _ := io.ReadAll(readerFile)
		//fmt.Println(string(data))

		readerFile, _ := file.Open()
		writerFile, _ := os.Create("./uploads/pzw.png")
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)
		fmt.Println(n)

		c.JSON(200, gin.H{"msg": "upload success"})
	})

	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("upload success %d files", len(files))})
	})

	router.Run("127.0.0.1:8088")
}
