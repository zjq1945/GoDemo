package main

import "github.com/gin-gonic/gin"
import "time"

func main() {
	r := gin.Default()
	r.GET("/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{"Current Time": time.Now()})
	})
	r.Run(":8008")
}
