package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "posted!")
	})

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8081")
}
