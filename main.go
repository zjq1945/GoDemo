package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	serviceHost := gin.Default()

	serviceHost.GET("/HelloWorld", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello World!": "jacky",
		})
	})

	serviceHost.Run(":8080")
}
