package restapi

import (
	"demo/db/mysql"
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func StartTest01Services() {
	serviceHost := gin.Default()

	serviceHost.GET("/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.Run(":8080")

}

func StartDemoServices() {
	serviceHost := gin.Default()

	serviceHost.GET("/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.GET("/GetConsumers", func(c *gin.Context) {
		consumers := jackymysql.GetConsumers()
		c.JSON(200, consumers)
	})

	serviceHost.Run(":8080")

}
