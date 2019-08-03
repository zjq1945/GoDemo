package restapi

import (
	"demo/db/mysql"
	// "encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func StartTest01Services() {
	serviceHost := gin.Default()

	serviceHost.GET("Demo/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.Run(":8080")

}

func StartDemoServices() {
	serviceHost := gin.Default()

	serviceHost.GET("/Demo/GetTimeNow", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Current Time": time.Now(),
		})
	})

	serviceHost.GET("/Consumer/GetConsumers", func(c *gin.Context) {
		consumers := jackymysql.GetConsumers()
		c.JSON(200, consumers)
	})

	serviceHost.GET("/Consumer/GetConsumer/:optional", func(c *gin.Context) {
		opt := c.Param("optional")
		name := strings.Replace(opt, "/", "", 1)

		consumer, result := jackymysql.GetConsumer(name)

		if result {
			c.JSON(200, consumer)
		} else {
			c.JSON(200, "")
		}
	})

	serviceHost.Run(":8080")

}
